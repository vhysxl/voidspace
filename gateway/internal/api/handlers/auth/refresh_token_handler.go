package auth

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	refreshToken := cookie.Value

	claims, err := utils.VerifyToken(refreshToken, h.PublicKey)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusUnauthorized, shared_constants.Unauthorized)
	}

	userID := claims["ID"].(string)
	username := claims["Username"].(string)

	res, err := h.UserService.RefreshToken(ctx, userID, username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to refresh token")
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
