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
	user := c.Get("authUser").(*models.AuthUser)

	p := new(models.PostRequest)
	err := c.Bind(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if p.Content == "" && len(p.PostImages) == 0 {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "Post must contain either content or images")
	}

	err = ph.Validator.Struct(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.Create(ctx, user.Username, user.ID, p)
	if err != nil {
		return utils.HandleDialError(ph.Logger, c, err, "failed to create post")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.PostCreated, res)
}

func (ph *PostHandler) GetPost(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	user, _ := val.(*models.AuthUser)
	if user == nil {
		user = &models.AuthUser{}
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	p := &models.GetPostRequest{ID: postId}
	if err := ph.Validator.Struct(p); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.GetPost(ctx, p, user.Username, user.ID)
	if err != nil {
		return utils.HandleDialError(ph.Logger, c, err, "failed to get post")

	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetPostSuccess, res)
}

func (ph *PostHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "invalid post id")
	}

	p := new(models.PostRequest)
	if err := c.Bind(p); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := ph.Validator.Struct(p); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if p.Content == "" && len(p.PostImages) == 0 {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrNoField)
	}

	if err := ph.PostService.Update(ctx, p, postId, user.Username, user.ID); err != nil {
		return utils.HandleDialError(ph.Logger, c, err, "failed to update post")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.UpdatePostSuccess, nil)
}

func (ph *PostHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := ph.PostService.Delete(ctx, postId, user.Username, user.ID); err != nil {
		return utils.HandleDialError(ph.Logger, c, err, "failed to delete post")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.DeletePostSuccess, nil)
}

func (ph *PostHandler) GetUserPosts(c echo.Context) error {
	ctx := c.Request().Context()
	username := c.Param("username")

	// val := c.Get("authUser")
	// user, _ := val.(*models.AuthUser)

	data := &models.GetUserRequest{Username: username}
	if err := ph.Validator.Struct(data); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ph.PostService.GetUserPosts(ctx, data.Username)
	if err != nil {
		return utils.HandleDialError(ph.Logger, c, err, "failed to get user posts")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetPostSuccess, res)
}
