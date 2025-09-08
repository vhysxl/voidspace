package handlers

import (
	"net/http"
	"strconv"
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

type PostHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	PostService    *service.PostService
}

func NewPostHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	postService *service.PostService,
) *PostHandler {
	return &PostHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		PostService:    postService,
	}
}

func (ph *PostHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	p := new(models.PostRequest)
	err := c.Bind(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ph.Validator.Struct(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.Create(ctx, username, ID, p)
	if err != nil {
		ph.Logger.Error("failed to create post", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.PostCreated, res)
}

func (ph *PostHandler) GetPost(c echo.Context) error {
	ctx := c.Request().Context()

	userID, _ := c.Get("ID").(string)
	username, _ := c.Get("username").(string)

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	p := &models.GetPostRequest{
		ID: postId,
	}

	err = ph.Validator.Struct(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.GetPost(ctx, p, username, userID)
	if err != nil {
		ph.Logger.Error("failed to Get post", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetPostSuccess, res)
}

func (ph *PostHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "invalid post id")
	}

	p := new(models.PostRequest)
	err = c.Bind(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ph.Validator.Struct(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	err = ph.PostService.Update(ctx, p, postId, username, ID)
	if err != nil {
		ph.Logger.Error("failed to update post", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdatePostSuccess, nil)
}

func (ph *PostHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ph.PostService.Delete(ctx, postId, username, ID)
	if err != nil {
		ph.Logger.Error("failed to Delete post", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.DeletePostSuccess, nil)
}

func (ph *PostHandler) GetUserPosts(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("username")

	data := &models.GetUserRequest{
		Username: username,
	}

	err := ph.Validator.Struct(data)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.GetUserPosts(ctx, data.Username)
	if err != nil {
		ph.Logger.Error("failed to Get posts", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetPostSuccess, res)
}
