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

	f := new(models.GetGlobalFeedReq)
	err := c.Bind(f)
	if err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	res, err := fh.FeedService.GetGlobalFeed(ctx, f)
	if err != nil {
		fh.Logger.Error("failed to fetch feed", zap.Error(err))
		code, msg := utils.GRPCErrorToHTTP(err)
		return responses.ErrorResponseMessage(c, code, msg)
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, "Get Feed Success", res)
}
