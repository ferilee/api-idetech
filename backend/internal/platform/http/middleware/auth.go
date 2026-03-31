package middleware

import (
	"context"
	"net/http"
	"strings"

	authservice "github.com/ferilee/api-idetech/backend/internal/auth/service"
)

type authClaimsContextKey string

const claimsKey authClaimsContextKey = "auth_claims"

func RequireAuth(authService *authservice.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := strings.TrimSpace(r.Header.Get("Authorization"))
			if header == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
			if token == "" || token == header {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			claims, err := authService.ParseToken(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), claimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ClaimsFromContext(ctx context.Context) (*authservice.TokenClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(*authservice.TokenClaims)
	return claims, ok
}
