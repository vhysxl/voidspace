package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserUsecase           usecase.UserUsecase
	Validator             *validator.Validate
	Logger                *zap.Logger
	HandlerContextTimeout time.Duration
}

func NewUserHandler(
	usecase usecase.UserUsecase,
	validator *validator.Validate,
	handlerTimeout time.Duration,
	logger *zap.Logger,
) *UserHandler {
	return &UserHandler{
		UserUsecase:           usecase,
		Validator:             validator,
		Logger:                logger,
		HandlerContextTimeout: handlerTimeout,
	}
}

func (uh *UserHandler) HandleGetCurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), uh.HandlerContextTimeout)
	defer cancel()

	userIDHeader := r.Header.Get("X-User-ID")
	if userIDHeader == "" {
		uh.Logger.Warn("Missing X-User-ID header")
		response.JSONErr(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	ID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		uh.Logger.Warn("Invalid X-User-ID header", zap.String("header", userIDHeader))
		response.JSONErr(w, http.StatusBadRequest, "Invalid User ID header")
		return
	}

	user, err := uh.UserUsecase.GetCurrentUser(ctx, ID)
	if err != nil {
		uh.Logger.Error("Usecase error", zap.Error(err))
		switch err {
		case ctx.Err():
			response.JSONErr(w, http.StatusRequestTimeout, ErrRequestTimeout)
			return
		case domain.ErrUserNotFound:
			response.JSONErr(w, http.StatusNotFound, err.Error())
			return
		default:
			response.JSONErr(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}
	}

	response.JSONSuccess(w, http.StatusOK, "Get current user success", map[string]any{
		"user": user,
	})
}
