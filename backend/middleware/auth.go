package middleware

import (
	"net/http"
	"os"
)

func writeNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Not allowed"))
}

func IsAllowed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyHeader := r.Header.Get("X-API-Key")
		if keyHeader == "" {
			writeNotAllowed(w)
			return
		}
		validApiKey := os.Getenv("API_KEY")
		if validApiKey == "" {
			writeNotAllowed(w)
			return
		}
		if keyHeader != validApiKey {
			writeNotAllowed(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
