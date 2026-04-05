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

func (h *PostHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	req := new(models.CreatePostRequest)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if err := h.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if req.Content == "" && len(req.PostImages) == 0 {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrNoField)
	}

	if err := h.PostService.Update(ctx, req, postID, user.Username, user.ID); err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to update post")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdatePostSuccess, nil)
}
