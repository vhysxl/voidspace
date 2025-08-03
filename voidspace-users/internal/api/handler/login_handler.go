package handler

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/token"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type LoginRequest struct {
	Credential string `json:"credential" validate:"required,min=3,max=50"`
	Password   string `json:"password" validate:"required,min=8"`
}

type LoginHandler struct {
	LoginUsecase          usecase.LoginUsecase
	Validator             *validator.Validate
	PrivateKey            *rsa.PrivateKey
	HandlerContextTimeout time.Duration
	AccessTokenDuration   time.Duration
	RefreshTokenDuration  time.Duration
	Logger                *zap.Logger
}

func NewLoginHandler(
	loginUsecase usecase.LoginUsecase,
	validator *validator.Validate,
	privateKey *rsa.PrivateKey,
	handlerTimeout time.Duration,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	logger *zap.Logger,
) *LoginHandler {
	return &LoginHandler{
		LoginUsecase:          loginUsecase,
		Validator:             validator,
		PrivateKey:            privateKey,
		HandlerContextTimeout: handlerTimeout,
		AccessTokenDuration:   accessTokenDuration,
		RefreshTokenDuration:  refreshTokenDuration,
		Logger:                logger,
	}
}

func (lh *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), lh.HandlerContextTimeout)
	defer cancel()
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lh.Logger.Error("Decoder error", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = lh.Validator.Struct(request)
	if err != nil {
		lh.Logger.Debug("Validation failed", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	user, err := lh.LoginUsecase.Login(ctx, request.Credential, request.Password)
	if err != nil {
		lh.Logger.Error("Usecase error", zap.Error(err))
		switch err {
		case ctx.Err():
			response.JSONErr(w, http.StatusRequestTimeout, ErrRequestTimeout)
			return
		case domain.ErrInvalidCredentials:
			response.JSONErr(w, http.StatusUnauthorized, err.Error())
			return
		default:
			response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}
	}

	accessToken, err := token.CreateAccessToken(user, lh.PrivateKey, lh.AccessTokenDuration)
	if err != nil {
		lh.Logger.Error("Generate access token error", zap.Error(err))
		response.JSONErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := token.CreateRefreshToken(user, lh.PrivateKey, lh.RefreshTokenDuration)
	if err != nil {
		lh.Logger.Error("Generate refresh token error", zap.Error(err))
		response.JSONErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   int(lh.RefreshTokenDuration.Seconds()),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	response.JSONSuccess(w, http.StatusCreated, "User logged in successfully", map[string]any{
		"access_token": accessToken,
		"expires_in":   int(lh.AccessTokenDuration.Seconds())})
}
