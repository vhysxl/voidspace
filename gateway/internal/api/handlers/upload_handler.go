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

type UploadHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UploadService  *service.UploadService
}

func NewUploadHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	uploadService *service.UploadService,
) *UploadHandler {
	return &UploadHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		UploadService:  uploadService,
	}
}

func (uh *UploadHandler) GenerateSignedURL(c echo.Context) error {
	var req models.SignedURLRequest

	// Bind JSON request
	if err := c.Bind(&req); err != nil {
		uh.Logger.Error("failed to bind", zap.Error(err))
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, err.Error())
	}

	// Validate request
	if err := uh.Validator.Struct(req); err != nil {
		uh.Logger.Error("failed to validate", zap.Error(err))
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, utils.FormatValidationError(err))
	}

	// Validate content type
	if !utils.IsValidImageType(req.ContentType) {
		uh.Logger.Error("invalid image type")
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "invalid filetype")
	}

	// Generate filename
	fileName := utils.GenerateUniqueFileName(req.ContentType)

	// Generate signed URL
	signedUrl, err := uh.UploadService.GenerateSignedURL(fileName, req.ContentType)
	if err != nil {
		uh.Logger.Error("failed generate signed url:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate signed URL"})
	}

	publicUrl := uh.UploadService.GetPublicURL(fileName)

	return responses.SuccessResponseMessage(c, http.StatusOK, "Url Successfully generated", map[string]any{
		"signedUrl": signedUrl,
		"publicUrl": publicUrl,
	})
}
