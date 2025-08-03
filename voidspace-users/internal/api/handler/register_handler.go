package handler

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"time"

	"net/http"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/token"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterHandler struct {
	RegisterUsecase       usecase.RegisterUsecase
	Validator             *validator.Validate
	PrivateKey            *rsa.PrivateKey
	HandlerContextTimeout time.Duration
	AccessTokenDuration   time.Duration
	RefreshTokenDuration  time.Duration
	Logger                *zap.Logger
}

func NewRegisterHandler(
	usecase usecase.RegisterUsecase,
	validator *validator.Validate,
	privateKey *rsa.PrivateKey,
	handlerTimeout time.Duration,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	logger *zap.Logger,
) *RegisterHandler {
	return &RegisterHandler{
		RegisterUsecase:       usecase,
		Validator:             validator,
		PrivateKey:            privateKey,
		HandlerContextTimeout: handlerTimeout,
		AccessTokenDuration:   accessTokenDuration,
		RefreshTokenDuration:  refreshTokenDuration,
		Logger:                logger,
	}
}

func (rh *RegisterHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), rh.HandlerContextTimeout)
	defer cancel()
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rh.Logger.Error("Decoder error", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = rh.Validator.Struct(request)
	if err != nil {
		rh.Logger.Debug("Validation failed", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		rh.Logger.Error("Hashing password error", zap.Error(err))
		response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	hashed := string(hashedPassword)

	user, err := rh.RegisterUsecase.Register(ctx, request.Username, request.Email, hashed)
	if err != nil {
		rh.Logger.Error("Usecase error", zap.Error(err))
		switch err {
		case ctx.Err():
			response.JSONErr(w, http.StatusRequestTimeout, ErrRequestTimeout)
			return
		case domain.ErrUserExists:
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		case domain.ErrEmailExists:
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		default:
			response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}
	}

	//goroutines next
	accessToken, err := token.CreateAccessToken(user, rh.PrivateKey, rh.AccessTokenDuration)
	if err != nil {
		rh.Logger.Error("Generate access token error", zap.Error(err))
		response.JSONErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := token.CreateRefreshToken(user, rh.PrivateKey, rh.RefreshTokenDuration)
	if err != nil {
		rh.Logger.Error("Generate refresh token error", zap.Error(err))
		response.JSONErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   int(rh.RefreshTokenDuration.Seconds()),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	response.JSONSuccess(w, http.StatusCreated, "User registered successfully", map[string]any{
		"access_token": accessToken,
		"expires_in":   int(rh.AccessTokenDuration.Seconds())})
}
