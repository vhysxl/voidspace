package service

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

type UploadService struct {
	Bucket string
	Client *storage.Client
}

// NewUploadService inisialisasi UploadService
func NewUploadService(ctx context.Context, bucket string) (*UploadService, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %w", err)
	}

	return &UploadService{
		Bucket: bucket,
		Client: client,
	}, nil
}

// Close release storage client
func (s *UploadService) Close() error {
	if s.Client != nil {
		return s.Client.Close()
	}
	return nil
}

// GenerateSignedURL generate signed URL untuk upload
func (s *UploadService) GenerateSignedURL(folder, fileName, contentType string) (string, error) {
	objectPath := fmt.Sprintf("%s/%s", folder, fileName)
	opts := &storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		Method:      "PUT",
		Expires:     time.Now().Add(15 * time.Minute),
		ContentType: contentType,
	}

	url, err := s.Client.Bucket(s.Bucket).SignedURL(objectPath, opts)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}

	return url, nil
}

// GetPublicURL generate public URL setelah upload
func (s *UploadService) GetPublicURL(folder, fileName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s/%s", s.Bucket, folder, fileName)
}
