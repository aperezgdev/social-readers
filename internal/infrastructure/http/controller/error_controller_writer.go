package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	domain_errors "github.com/aperezgdev/social-readers-api/internal/domain/errors"
)

type validationErrorResponse struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, err error) {
	if validationError, ok := err.(domain_errors.ValidationError); ok {
		w.WriteHeader(http.StatusBadRequest)
		errEncoding := json.NewEncoder(w).Encode(validationErrorResponse{
			Field: validationError.Field,
			Message: err.Error(),
		})

		if errEncoding != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if errors.Is(err, domain_errors.ErrNotExistPost) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if errors.Is(err, domain_errors.ErrNotExistUser) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}