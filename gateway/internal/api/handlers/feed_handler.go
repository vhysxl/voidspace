package handlers

import (
	"net/http"
	"strconv"
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

	userID, _ := c.Get("ID").(string)
	username, _ := c.Get("username").(string)

	cursor := c.QueryParam("cursor")
	cursorID := c.QueryParam("cursorid")

	var cursorTime time.Time
	var cursorIDInt int

	// Convert cursor string to time.Time
	if cursor != "" {
		if parsedTime, err := time.Parse(time.RFC3339, cursor); err == nil {
			cursorTime = parsedTime
		} else {
			// Jika gagal parse RFC3339, coba parse sebagai Unix timestamp
			if timestamp, err := strconv.ParseInt(cursor, 10, 64); err == nil {
				parsedTime := time.Unix(timestamp, 0)
				cursorTime = parsedTime
			}
		}
	}

	if cursorID != "" {
		if id, err := strconv.Atoi(cursorID); err == nil {
			cursorIDInt = id
		}
	}

	f := &models.GetGlobalFeedReq{
		Cursor:   cursorTime,
		CursorID: cursorIDInt,
	}

	res, err := fh.FeedService.GetGlobalFeed(ctx, f, username, userID)
	if err != nil {
		fh.Logger.Error("failed to fetch feed", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, "Get Feed Success", res)
}

func (fh *FeedHandler) GetFollowFeed(c echo.Context) error {
	ctx := c.Request().Context()

	userID, _ := c.Get("ID").(string)
	username, _ := c.Get("username").(string)

	cursor := c.QueryParam("cursor")
	cursorID := c.QueryParam("cursorid")

	var cursorTime time.Time
	var cursorIDInt int

	// Convert cursor string to time.Time
	if cursor != "" {
		if parsedTime, err := time.Parse(time.RFC3339, cursor); err == nil {
			cursorTime = parsedTime
		} else {
			// Jika gagal parse RFC3339, coba parse sebagai Unix timestamp
			if timestamp, err := strconv.ParseInt(cursor, 10, 64); err == nil {
				parsedTime := time.Unix(timestamp, 0)
				cursorTime = parsedTime
			}
		}
	}

	if cursorID != "" {
		if id, err := strconv.Atoi(cursorID); err == nil {
			cursorIDInt = id
		}
	}

	f := &models.GetFollowFeedReq{
		Cursor:   cursorTime,
		CursorID: cursorIDInt,
	}

	res, err := fh.FeedService.GetFollowFeed(ctx, username, userID, f)
	if err != nil {
		fh.Logger.Error("failed to fetch feed", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, "Get Feed Success", res)
}
