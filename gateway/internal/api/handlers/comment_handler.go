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

type CommentHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	CommentService *service.CommentsService
}

func NewCommentHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	commentService *service.CommentsService,

) *CommentHandler {
	return &CommentHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		CommentService: commentService,
	}
}

func (ch *CommentHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	p := new(models.CreateCommentRequest)
	if err := c.Bind(p); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := ch.Validator.Struct(p); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ch.CommentService.Create(ctx, p, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(ch.Logger, c, err, "failed to create comment")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.CommentSuccess, res)
}

func (ch *CommentHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ch.CommentService.Delete(ctx, int32(commentId), user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(ch.Logger, c, err, "failed to delete comment")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.CommentDeleteSuccess, nil)
}

func (ch *CommentHandler) GetAllByPostID(c echo.Context) error {
	ctx := c.Request().Context()

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	res, err := ch.CommentService.GetAllByPostID(ctx, int32(postId))
	if err != nil {
		return utils.HandleDialError(ch.Logger, c, err, "failed to get comments by post ID")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetCommentsSuccess, res)
}

func (ch *CommentHandler) GetAllByUser(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.Param("username")
	if username == "" {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	res, err := ch.CommentService.GetAllByUser(ctx, username)
	if err != nil {
		return utils.HandleDialError(ch.Logger, c, err, "failed to get comments by user ID")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetCommentsSuccess, res)
}
