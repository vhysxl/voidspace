package handler

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"log"
	"time"

	"net/http"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/token"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
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
}

func NewRegisterHandler(
	usecase usecase.RegisterUsecase,
	validator *validator.Validate,
	privateKey *rsa.PrivateKey,
	handlerTimeout time.Duration,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *RegisterHandler {
	return &RegisterHandler{
		RegisterUsecase:       usecase,
		Validator:             validator,
		PrivateKey:            privateKey,
		HandlerContextTimeout: handlerTimeout,
		AccessTokenDuration:   accessTokenDuration,
		RefreshTokenDuration:  refreshTokenDuration,
	}
}

func (rh *RegisterHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), rh.HandlerContextTimeout)
	defer cancel()
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("[RegisterHandler] Failed to decode request body: %v", err)
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = rh.Validator.Struct(request)
	if err != nil {
		log.Printf("[RegisterHandler] Validation failed: %v", err)
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Printf("[RegisterHandler] Error hashing password: %v", err)
		response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	hashed := string(hashedPassword)

	user, err := rh.RegisterUsecase.Register(ctx, request.Username, request.Email, hashed)
	if err != nil {
		log.Printf("[RegisterHandler] Usecase.Register error: %v", err)
		switch err {
		case domain.ErrUserExists:
			log.Printf("[RegisterHandler] User exists error: %v", err)
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		case domain.ErrEmailExists:
			log.Printf("[RegisterHandler] Email exists error: %v", err)
			response.JSONErr(w, http.StatusConflict, err.Error())
			return
		default:
			log.Printf("[RegisterHandler] Unexpected error in usecase.Register: %v", err)
			response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}
	}

	//goroutines next
	accessToken, err := token.CreateAccessToken(user, rh.PrivateKey, rh.AccessTokenDuration)
	if err != nil {
		log.Printf("[RegisterHandler] Error generating access token: %v", err)
		response.JSONErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := token.CreateRefreshToken(user, rh.PrivateKey, rh.RefreshTokenDuration)
	if err != nil {
		log.Printf("[RegisterHandler] Error generating refresh token: %v", err)
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

	response.JSONSuccess(w, http.StatusCreated, "User registered successfully", map[string]interface{}{
		"access_token": accessToken,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"expires_in": int(rh.AccessTokenDuration.Seconds())})
}
