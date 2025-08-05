package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	"voidspace/users/utils/response"
	"voidspace/users/utils/validations"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type GetUserParam struct {
	Username string `validate:"required,min=3,max=32,alphanum"`
}

// UpdateUserReq represents a partial update payload for updating user profile.
// All fields are optional and will only be updated if provided in the request.
// Use pointers to differentiate between "field not provided" (nil) and "empty value" ("").
type UpdateUserReq struct {
	DisplayName *string `json:"display_name" validate:"omitempty,max=50"`
	Bio         *string `json:"bio" validate:"omitempty,max=160"`
	AvatarUrl   *string `json:"avatar_url" validate:"omitempty,url"`
	BannerUrl   *string `json:"banner_url" validate:"omitempty,url"`
	Location    *string `json:"location" validate:"omitempty,max=100"`
}

type UserHandler struct {
	UserUsecase           usecase.UserUsecase
	ProfileUsecase        usecase.ProfileUsecase
	Validator             *validator.Validate
	Logger                *zap.Logger
	HandlerContextTimeout time.Duration
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
	profileUsecase usecase.ProfileUsecase,
	validator *validator.Validate,
	handlerTimeout time.Duration,
	logger *zap.Logger,
) *UserHandler {
	return &UserHandler{
		UserUsecase:           userUsecase,
		ProfileUsecase:        profileUsecase,
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
		response.JSONErr(w, http.StatusUnauthorized, ErrUnauthorized)
		return
	}

	ID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		uh.Logger.Warn("Invalid X-User-ID header", zap.String("header", userIDHeader))
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
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

func (uh *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), uh.HandlerContextTimeout)
	defer cancel()

	usernameParam := chi.URLParam(r, "username")

	if usernameParam == "" {
		uh.Logger.Warn("Missing username path parameter")
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	param := GetUserParam{
		Username: usernameParam,
	}

	err := uh.Validator.Struct(param)
	if err != nil {
		uh.Logger.Warn("Validation failed", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	user, err := uh.UserUsecase.GetUser(ctx, usernameParam)
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

	response.JSONSuccess(w, http.StatusOK, "Get user success", map[string]any{
		"user": user,
	})
}

func (uh *UserHandler) HandleUpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), uh.HandlerContextTimeout)
	defer cancel()

	userIDHeader := r.Header.Get("X-User-ID")
	if userIDHeader == "" {
		uh.Logger.Warn("Missing X-User-ID header")
		response.JSONErr(w, http.StatusUnauthorized, ErrUnauthorized)
		return
	}

	ID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		uh.Logger.Warn("Invalid X-User-ID header", zap.String("header", userIDHeader))
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	var request UpdateUserReq

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		uh.Logger.Error("Decoder error", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	err = uh.Validator.Struct(request)
	if err != nil {
		uh.Logger.Warn("Validation failed", zap.Error(err))
		response.JSONErr(w, http.StatusBadRequest, validations.FormatValidationError(err))
		return
	}

	profile := &domain.Profile{
		UserID:      ID,
		DisplayName: request.DisplayName,
		Bio:         request.Bio,
		AvatarUrl:   request.AvatarUrl,
		BannerUrl:   request.BannerUrl,
		Location:    request.Location,
	}

	err = uh.ProfileUsecase.UpdateProfile(ctx, ID, profile)
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
}
