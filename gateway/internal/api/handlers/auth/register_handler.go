package auth

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	r := new(models.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if err := h.Validator.Struct(r); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := h.UserService.Register(ctx, r)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to register")
	}

	if res.RefreshToken != "" {
		err := utils.SetRefreshTokenCookie(c, res.RefreshToken)
		if err != nil {
			h.Logger.Error("Failed to set refresh token cookie", zap.Error(err))
			return utils.HandleDialError(h.Logger, c, err, "failed to set refresh token cookie")
		}
	}

	registerRes := models.RegisterResponseAPI{
		AccessToken: res.AccessToken,
		ExpiresIn:   res.ExpiresIn,
	}

	return responses.SuccessResponseMessage(c,
		http.StatusCreated,
		constants.RegisterSuccess,
		registerRes,
	)
}
