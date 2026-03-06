package common

import (
	"fmt"
	"strings"
)

func GenerateDBPlaceholders[T any](data []T) (string, []any) {
	placeholders := make([]string, len(data))
	args := make([]any, len(data))
	for i, v := range data {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = v
	}

	finalQuery := strings.Join(placeholders, ",")

	return finalQuery, args

}
