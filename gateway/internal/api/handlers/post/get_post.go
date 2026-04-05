package post

import (
	"net/http"
	"strconv"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *PostHandler) GetPost(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	authUser, _ := val.(*models.AuthUser)
	if authUser == nil {
		authUser = &models.AuthUser{}
	}

	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	res, err := h.PostService.GetPost(ctx, postID, authUser.Username, authUser.ID)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to get post")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetPostSuccess, res)
}
