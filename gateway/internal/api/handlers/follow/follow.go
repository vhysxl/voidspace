package follow

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *FollowHandler) Follow(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)
	targetUsername := c.Param("username")

	if targetUsername == "" {
		return responses.ErrorResponseMessage(
			c,
			http.StatusBadRequest,
			constants.ErrUsernameRequired,
		)
	}

	if len(targetUsername) > 50 {
		return responses.ErrorResponseMessage(
			c,
			http.StatusBadRequest,
			shared_constants.InvalidRequest,
		)
	}

	err := h.UserService.Follow(ctx, user.ID, user.Username, targetUsername)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to follow user")
	}

	return responses.SuccessResponseMessage(
		c, http.StatusOK,
		constants.FollowSuccess,
		nil,
	)
}
