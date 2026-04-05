package post

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"

	"github.com/labstack/echo/v4"
)

func (h *PostHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	req := new(models.CreatePostRequest)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if req.Content == "" && len(req.PostImages) == 0 {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "Post must contain either content or images")
	}

	if err := h.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := h.PostService.Create(ctx, user.Username, user.ID, req)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to create post")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.PostCreated, res)
}
