package post

import (
	"net/http"
	"strconv"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *PostHandler) LikePost(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "invalid post id")
	}

	if err := h.PostService.LikePost(ctx, postID, user.Username, user.ID); err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to like post")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.LikeSuccess, nil)
}
