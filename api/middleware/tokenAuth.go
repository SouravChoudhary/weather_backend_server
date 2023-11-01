package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"weatherapp/pkg/auth"
)

// TokenAuthMiddleware is middleware for token-based authentication.
func TokenAuthMiddleware(jwtAuth auth.JWTAuthenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenString, err := extractTokenFromHeader(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			userID, username, err := verifyToken(tokenString, jwtAuth)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Store user information in the request context
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "username", username)

			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("Invalid Authorization header format")
	}

	return parts[1], nil
}

// verifyToken verifies a JWT token and returns the claims.
func verifyToken(tokenString string, jwtAuth auth.JWTAuthenticator) (int, string, error) {
	claims, err := jwtAuth.ValidateToken(tokenString)
	if err != nil {
		return 0, "", err
	}
	return claims.UserID, claims.Username, nil
}
