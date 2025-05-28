package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"io"

	"books-api/app/request"
)

type key string

const CreateBookKey key = "createBookReq"

func RequestValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.CreateBookRequest

		// Access body from request and store manually to body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// handling custom validation
		if err := req.Validate(); err != nil {
			if ve, ok := err.(*request.ValidationError); ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"message": "Validation failed",
					"errors":  ve.Errors,
				})
				return
			}

			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		// End custom validation

		// Store the validated request in context
		ctx := context.WithValue(r.Context(), CreateBookKey, req)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}