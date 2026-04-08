package models

type SignedURLRequest struct {
	ContentType string `json:"contentType" validate:"required"`
	Folder      string `json:"folder" validate:"required,oneof=posts avatars banners"`
}
