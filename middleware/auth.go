package middleware

import (
	"log"
	"net/http"
)

func ValidateAPIKey(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check for API key in the header
			key := r.Header.Get("api_key")
			log.Printf("key: %s", key)
			if key == "" || key != apiKey {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
