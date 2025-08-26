package handlers

import (
	"crypto/rsa"
	"net/http"
	"time"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/internal/service"

	"voidspaceGateway/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in,omitempty"`
}

type AuthHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	AuthService    *service.AuthService
	PublicKey      *rsa.PublicKey
}

func NewAuthHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	authService *service.AuthService,
	PublicKey *rsa.PublicKey,
) *AuthHandler {
	return &AuthHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		AuthService:    authService,
		PublicKey:      PublicKey,
	}
}

func (ah *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	l := new(models.LoginRequest)
	err := c.Bind(l)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ah.Validator.Struct(l)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ah.AuthService.Login(ctx, l)
	if err != nil {
		ah.Logger.Error("failed to login", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	if res.RefreshToken != "" {
		utils.SetRefreshTokenCookie(c, res.RefreshToken)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.LoginSuccess, &response{AccessToken: res.AccessToken, ExpiresIn: res.ExpiresIn})
}

func (ah *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	r := new(models.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := ah.Validator.Struct(r); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ah.AuthService.Register(ctx, r)
	if err != nil {
		ah.Logger.Error("failed to register", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	if res.RefreshToken != "" {
		utils.SetRefreshTokenCookie(c, res.RefreshToken)
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.RegisterSuccess, &response{AccessToken: res.AccessToken, ExpiresIn: res.ExpiresIn})
}

func (ah *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	refreshToken := cookie.Value

	claims, err := utils.VerifyRefreshToken(refreshToken, ah.PublicKey)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
	}

	userID := claims["ID"].(string)
	username := claims["Username"].(string)

	res, err := ah.AuthService.RefreshToken(ctx, userID, username)
	if err != nil {
		ah.Logger.Error("failed to refresh token", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	if res.RefreshToken != "" {
		utils.SetRefreshTokenCookie(c, res.RefreshToken)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.TokenRefresh, &response{AccessToken: res.AccessToken, ExpiresIn: res.ExpiresIn})
}

func (ah *AuthHandler) Logout(c echo.Context) error {
	// Clear refresh token cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   false, // set to true in production
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	}
	c.SetCookie(cookie)

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.LogoutSuccess, nil)
}
