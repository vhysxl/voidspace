package follow

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *FollowHandler) Follow(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	requestBody := new(models.FollowRequest)
	if err := c.Bind(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if err := h.Validator.Struct(requestBody); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err := h.UserService.Follow(ctx, user.ID, user.Username, requestBody.TargetUsername)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to follow user")
	}

	return responses.SuccessResponseMessage(
		c, http.StatusOK,
		constants.FollowSuccess,
		nil,
	)
}
