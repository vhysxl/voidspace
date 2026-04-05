package comment

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *CommentHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	req := new(models.CreateCommentRequest)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if err := h.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err := h.CommentService.Create(ctx, req, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to create comment")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.CommentSuccess, nil)
}
