package handlers

import (
	"net/http"
	"time"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/internal/service"
	"voidspaceGateway/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type FeedHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	FeedService    *service.FeedService
}

func NewFeedHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	FeedService *service.FeedService,
) *FeedHandler {
	return &FeedHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		FeedService:    FeedService,
	}
}

func (fh *FeedHandler) GetGlobalFeed(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	user, _ := val.(*models.AuthUser)
	if user == nil {
		user = &models.AuthUser{}
	}

	cursor := c.QueryParam("cursor")
	cursorID := c.QueryParam("cursorid")

	cursorTime, cursorIDInt := utils.ExtractCursor(cursor, cursorID)

	f := &models.GetGlobalFeedRequest{
		Cursor:   cursorTime,
		CursorID: cursorIDInt,
	}

	res, err := fh.FeedService.GetGlobalFeed(ctx, f, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(fh.Logger, c, err, "failed to fetch feed")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, "Get Feed Success", res)
}

func (fh *FeedHandler) GetFollowFeed(c echo.Context) error {
	ctx := c.Request().Context()

	val := c.Get("authUser")
	user, _ := val.(*models.AuthUser)
	if user == nil {
		user = &models.AuthUser{}
	}

	cursor := c.QueryParam("cursor")
	cursorID := c.QueryParam("cursorid")

	cursorTime, cursorIDInt := utils.ExtractCursor(cursor, cursorID)

	f := &models.GetFollowFeedRequest{
		Cursor:   cursorTime,
		CursorID: cursorIDInt,
	}

	res, err := fh.FeedService.GetFollowFeed(ctx, user.ID, user.Username, f)
	if err != nil {
		return utils.HandleDialError(fh.Logger, c, err, "failed to fetch follow feed")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, "Get Feed Success", res)
}
