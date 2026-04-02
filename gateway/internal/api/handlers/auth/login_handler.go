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

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	requestBody := new(models.LoginRequest)
	if err := c.Bind(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := h.Validator.Struct(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := h.UserService.Login(ctx, requestBody)
	if err != nil {
		h.Logger.Error("failed to login", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	if res.RefreshToken != "" {
		err := utils.SetRefreshTokenCookie(c, res.RefreshToken)
		if err != nil {
			h.Logger.Error("Failed to set refresh token cookie", zap.Error(err))
			return utils.HandleDialError(h.Logger, c, err, "failed to set refresh token cookie")
		}
	}

	loginRes := models.LoginResponseAPI{
		AccessToken: res.AccessToken,
		ExpiresIn:   res.ExpiresIn,
	}

	return responses.SuccessResponseMessage(
		c, http.StatusOK,
		constants.LoginSuccess,
		loginRes,
	)
}
