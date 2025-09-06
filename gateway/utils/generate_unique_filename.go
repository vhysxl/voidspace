package utils

import (
	"fmt"
	"time"
)

func GenerateUniqueFileName(contentType string) string {
	ext := GetFileExtension(contentType)
	timestamp := time.Now().UnixNano() //biar bisa beda
	return fmt.Sprintf("%d%s", timestamp, ext)
}
