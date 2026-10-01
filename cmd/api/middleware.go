package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/droffilc1/webhook-service/internal/auth"
	"github.com/droffilc1/webhook-service/internal/store"
)

func apiKeyMiddleware(st store.Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		rawKey := parts[1]
		hash := auth.Hash(rawKey)

		apiKey, err := st.GetAPIKeyByHash(hash)
		if err != nil {
			http.Error(w, "invalid API key", http.StatusUnauthorized)
			return
		}

		if apiKey.RevokedAt != nil {
			http.Error(w, "API key has been revoked", http.StatusUnauthorized)
			return
		}

		if apiKey.ExpiresAt != nil && time.Now().After(*apiKey.ExpiresAt) {
			http.Error(w, "API key has expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		h.Set("Content-Security-Policy", "default-src 'self'")
		h.Set("X-Frame-Options", "DENY")
		h.Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
