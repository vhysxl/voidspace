package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
	Data    T      `json:"data"`
}

func JSONSuccess[T any](w http.ResponseWriter, statuscode int, detail string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(SuccessResponse[T]{
		Success: true,
		Detail:  detail,
		Data:    data,
	})
}
