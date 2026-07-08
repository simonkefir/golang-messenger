package core_http_middleware

import (
	"context"
	"net/http"
	"strings"

	core_jwt "github.com/simonkefir/golang-messenger/internal/core/jwt"
)

type contextKey string

const (
	userIDKey contextKey = "user id"
)

func setUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(userIDKey).(int64)
	return userID, ok
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error":"missing authorization header"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"error":"invalid authorization header"}`, http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		userID, err := core_jwt.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusUnauthorized)
			return
		}

		ctx := setUserID(r.Context(), userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
