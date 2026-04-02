package auth

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	refreshToken := cookie.Value

	claims, err := utils.VerifyRefreshToken(refreshToken, h.PublicKey)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
	}

	userID := claims["ID"].(string)
	username := claims["Username"].(string)

	res, err := h.UserService.RefreshToken(ctx, userID, username)
	if err != nil {
		h.Logger.Error("failed to refresh token", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	refreshRes := models.RefreshTokenResponseAPI{
		AccessToken: res.AccessToken,
		ExpiresIn:   res.ExpiresIn,
	}

	return responses.SuccessResponseMessage(
		c,
		http.StatusOK,
		constants.TokenRefresh,
		refreshRes,
	)
}
