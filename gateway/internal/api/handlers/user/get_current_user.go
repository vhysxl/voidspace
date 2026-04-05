package user

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	res, err := h.UserService.GetCurrentUser(ctx, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to get current user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetProfileSuccess, res)
}
