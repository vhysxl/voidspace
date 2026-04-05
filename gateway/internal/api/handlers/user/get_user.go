package user

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	val := c.Get("authUser")
	authUser, _ := val.(*models.AuthUser)
	if authUser == nil {
		authUser = &models.AuthUser{}
	}
	username := c.Param("username")

	if username == "" {
		return responses.ErrorResponseMessage(
			c,
			http.StatusBadRequest,
			constants.ErrUsernameRequired,
		)
	}

	if len(username) > 50 {
		return responses.ErrorResponseMessage(
			c,
			http.StatusBadRequest,
			shared_constants.InvalidRequest,
		)
	}

	user, err := h.UserService.GetUser(ctx, username, authUser.ID, authUser.Username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to get user")
	}

	return responses.SuccessResponseMessage(
		c,
		http.StatusOK,
		constants.GetUserSuccess,
		user,
	)
}
