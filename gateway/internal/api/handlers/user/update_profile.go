package user

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	req := new(models.UpdateProfileRequest)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, shared_constants.InvalidRequest)
	}

	if *req == (models.UpdateProfileRequest{}) {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrNoField)
	}

	if err := h.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if err := h.UserService.UpdateProfile(ctx, user.ID, user.Username, req); err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to update profile")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdateProfileSuccess, nil)
}
