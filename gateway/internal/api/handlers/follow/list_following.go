package follow

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *FollowHandler) ListFollowing(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	authUser, _ := val.(*models.AuthUser)
	if authUser == nil {
		authUser = &models.AuthUser{}
	}

	username := c.Param("username")

	// Resolve username to user ID
	targetUser, err := h.UserService.GetUser(ctx, username, authUser.ID, authUser.Username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to resolve user")
	}

	following, err := h.UserService.ListFollowing(ctx, int64(targetUser.ID))
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to list following")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.ListFollowingSuccess, following)
}
