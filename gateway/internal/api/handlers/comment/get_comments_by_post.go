package comment

import (
	"net/http"
	"strconv"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *CommentHandler) GetAllByPostID(c echo.Context) error {
	ctx := c.Request().Context()

	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	res, err := h.CommentService.GetAllByPostID(ctx, postID)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to get comments by post ID")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetCommentsSuccess, res)
}
