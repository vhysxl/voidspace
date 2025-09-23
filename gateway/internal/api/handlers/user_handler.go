package handlers

import (
	"net/http"
	"time"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/internal/service"
	"voidspaceGateway/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UserService    *service.UserService
}

func NewUserHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	userService *service.UserService,
) *UserHandler {
	return &UserHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		UserService:    userService,
	}
}

func (uh *UserHandler) GetCurrentUser(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	res, err := uh.UserService.GetCurrentUser(ctx, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to get current user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetProfileSuccess, res)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	user, _ := val.(*models.AuthUser)
	if user == nil {
		user = &models.AuthUser{}
	}

	username := c.Param("username")
	req := &models.GetUserRequest{Username: username}

	if err := uh.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := uh.UserService.GetUser(ctx, req.Username, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to get user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetUserSuccess, res)
}

func (uh *UserHandler) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	req := new(models.UpdateProfileRequest)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if *req == (models.UpdateProfileRequest{}) {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrNoField)
	}

	if err := uh.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if err := uh.UserService.UpdateProfile(ctx, user.ID, user.Username, req); err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to update profile")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdateProfileSuccess, nil)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	if err := uh.UserService.DeleteUser(ctx, user.ID, user.Username); err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to delete user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.DeleteUserSuccess, nil)
}

func (uh *UserHandler) Follow(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	req := &models.FollowRequest{TargetUsername: c.Param("username")}
	if err := uh.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if err := uh.UserService.Follow(ctx, user.ID, user.Username, req.TargetUsername); err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to follow user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.FollowSuccess, nil)
}

func (uh *UserHandler) Unfollow(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	req := &models.FollowRequest{TargetUsername: c.Param("username")}
	if err := uh.Validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if err := uh.UserService.Unfollow(ctx, user.ID, user.Username, req.TargetUsername); err != nil {
		return utils.HandleDialError(uh.Logger, c, err, "failed to unfollow user")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UnfollowSuccess, nil)
}
