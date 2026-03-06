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

type LikeHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	LikeService    *service.LikeService
}

func NewLikeHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	Likeservice *service.LikeService,
) *LikeHandler {
	return &LikeHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		LikeService:    Likeservice,
	}
}

func (lh *LikeHandler) Like(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	l := &models.LikeRequest{
		PostID:   postId,
		UserID:   user.ID,
		Username: user.Username,
	}

	err = lh.Validator.Struct(l)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := lh.LikeService.Like(ctx, l)
	if err != nil {
		return utils.HandleDialError(lh.Logger, c, err, "Failed to like post")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.LikeSuccess, res)
}

func (lh *LikeHandler) Unlike(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("authUser").(*models.AuthUser)

	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	l := &models.LikeRequest{
		PostID:   postId,
		UserID:   user.ID,
		Username: user.Username,
	}

	err = lh.Validator.Struct(l)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	res, err := lh.LikeService.Unlike(ctx, l)
	if err != nil {
		return utils.HandleDialError(lh.Logger, c, err, "Failed to unlike post")
	}

	return responses.SuccessResponseMessage(c, http.StatusCreated, constants.UnlikeSuccess, res)
}
