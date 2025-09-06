package utils

import "slices"

func IsValidImageType(contentType string) bool {
	validTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}
	return slices.Contains(validTypes, contentType)
}
