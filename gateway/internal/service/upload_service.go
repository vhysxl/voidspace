package service

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

type UploadService struct {
	Bucket         string
	Client         *storage.Client
	CredentialPath string
}

// TODO: UPDATE THIS TO USE POLICY BECAUSE THERE IS NO SIZE VALIDATION IN SIGNED URL
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
func (s *UploadService) GenerateSignedURL(fileName, contentType string) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		Method:      "PUT",
		Expires:     time.Now().Add(1 * time.Millisecond), // FOR SECURITY REASON, CHANGE THIS LATER IN PROD
		ContentType: contentType,
	}

	url, err := s.Client.Bucket(s.Bucket).SignedURL(fileName, opts)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}

	return url, nil
}

// GetPublicURL generate public URL setelah upload
func (s *UploadService) GetPublicURL(fileName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.Bucket, fileName)
}
