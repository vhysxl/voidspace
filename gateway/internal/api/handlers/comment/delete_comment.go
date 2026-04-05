package comment

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

func (h *CommentHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	commentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if err := h.CommentService.Delete(ctx, user.Username, user.ID, commentID); err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to delete comment")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.CommentDeleteSuccess, nil)
}
