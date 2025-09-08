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

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	data := &models.GetProfileRequest{
		ID:       ID,
		Username: username,
	}

	err := uh.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	user, err := uh.UserService.GetCurrentUser(ctx, data.ID, data.Username)
	if err != nil {
		uh.Logger.Error("failed to get current user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}
	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetProfileSuccess, user)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("username")
	userID, _ := c.Get("ID").(string)
	usernameRequester, _ := c.Get("username").(string)

	u := &models.GetUserRequest{
		Username: username,
	}

	err := uh.Validator.Struct(u)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	user, err := uh.UserService.GetUser(ctx, username, userID, usernameRequester)
	if err != nil {
		uh.Logger.Error("failed to get user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetUserSuccess, user)
}

func (uh *UserHandler) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	data := &models.GetProfileRequest{
		ID:       ID,
		Username: username,
	}

	err := uh.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	r := new(models.UpdateProfileRequest)
	err = c.Bind(r)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = uh.Validator.Struct(r)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err = uh.UserService.UpdateProfile(ctx, data.ID, data.Username, r)
	if err != nil {
		uh.Logger.Error("failed to update profile", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdateProfileSuccess, nil)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	data := &models.GetProfileRequest{
		ID:       ID,
		Username: username,
	}

	err := uh.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err = uh.UserService.DeleteUser(ctx, data.ID, data.Username)
	if err != nil {
		uh.Logger.Error("failed to delete user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.DeleteUserSuccess, nil)
}

func (uh *UserHandler) Follow(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)
	targetUsername := c.Param("username")

	data := &models.GetProfileRequest{
		ID:       ID,
		Username: username,
	}

	err := uh.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	target := &models.FollowRequest{
		TargetUsername: targetUsername,
	}

	err = uh.Validator.Struct(target)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err = uh.UserService.Follow(ctx, data.ID, data.Username, target.TargetUsername)
	if err != nil {
		uh.Logger.Error("failed to follow user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.FollowSuccess, nil)
}

func (uh *UserHandler) Unfollow(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)
	targetUsername := c.Param("username")

	data := &models.GetProfileRequest{
		ID:       ID,
		Username: username,
	}

	err := uh.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	target := &models.FollowRequest{
		TargetUsername: targetUsername,
	}
	err = uh.Validator.Struct(target)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err = uh.UserService.Unfollow(ctx, data.ID, data.Username, target.TargetUsername)
	if err != nil {
		uh.Logger.Error("failed to unfollow user", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UnfollowSuccess, nil)
}
