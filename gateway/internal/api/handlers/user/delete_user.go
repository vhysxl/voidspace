package user

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	if err := h.UserService.DeleteUser(ctx, user.ID, user.Username); err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to delete user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.DeleteUserSuccess, nil)
}
