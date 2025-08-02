package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

func JSONErr(w http.ResponseWriter, statuscode int, detail string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Detail:  detail,
	})
}
