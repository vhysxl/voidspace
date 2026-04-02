package follow

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *FollowHandler) Follow(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	requestBody := new(models.FollowRequest)
	if err := c.Bind(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := h.Validator.Struct(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err := h.UserService.Follow(ctx, user.ID, user.Username, requestBody.TargetUsername)
	if err != nil {
		h.Logger.Error("failed to follow user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(
		c, http.StatusOK,
		constants.FollowSuccess,
		nil,
	)
}
