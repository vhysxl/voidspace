package models

type SignedURLRequest struct {
	ContentType string `json:"contentType" validate:"required"`
}
