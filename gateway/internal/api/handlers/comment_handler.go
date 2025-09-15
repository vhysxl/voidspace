package handlers

import (
	"fmt"
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

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	p := new(models.CreateCommentReq)
	fmt.Println(p)
	err := c.Bind(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ch.Validator.Struct(p)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := ch.CommentService.Create(ctx, p, ID, username)
	if err != nil {
		ch.Logger.Error("failed to create comment", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.CommentSuccess, res)
}

func (ch *CommentHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	ID := c.Get("ID").(string)
	username := c.Get("username").(string)

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	err = ch.CommentService.Delete(ctx, int32(commentId), ID, username)
	if err != nil {
		ch.Logger.Error("failed to delete comment", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.CommentDeleteSuccess, nil)
}

func (ch *CommentHandler) GetAllByPostID(c echo.Context) error {
	ctx := c.Request().Context()

	// Get postID from URL parameter
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	res, err := ch.CommentService.GetAllByPostID(ctx, int32(postId))
	if err != nil {
		ch.Logger.Error("failed to get comments by post ID", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
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
		ch.Logger.Error("failed to get comments by user ID", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetCommentsSuccess, res)
}
