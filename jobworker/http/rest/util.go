package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lht102/message-playground/jobworker/api"
)

func decode(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}

func respond(w http.ResponseWriter, statCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statCode)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func respondErr(w http.ResponseWriter, statusCode int, message string) {
	respond(w, statusCode, api.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	})
}
