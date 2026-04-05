package comment

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *CommentHandler) GetAllByUser(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("username")
	if username == "" {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	res, err := h.CommentService.GetAllByUser(ctx, username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to get comments by user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetCommentsSuccess, res)
}
