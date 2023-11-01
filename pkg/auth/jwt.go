package auth

import (
	"errors"
	"time"
	"weatherapp/config"
	"weatherapp/pkg/logging"

	"github.com/dgrijalva/jwt-go"
)

// JWTAuthenticator is an interface for JWT authentication.
type JWTAuthenticator interface {
	GenerateToken(userID int, username string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

// jwtAuthenticator is an implementation of JWTAuthenticator.
type jwtAuthenticator struct {
	secretKey []byte
	tokenTTL  time.Duration
	logger    logging.Logger
}

// Claims represents the JWT claims.
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewJWTAuthenticator creates a new JWTAuthenticator with the specified secret key and token TTL.
func NewJWTAuthenticator(config config.JWTAuthenticator, logger logging.Logger) JWTAuthenticator {
	return &jwtAuthenticator{
		secretKey: []byte(config.SecretKey),
		tokenTTL:  config.TokenTTL,
		logger:    logger,
	}
}

// GenerateToken generates a JWT token with the user ID and username.
func (a *jwtAuthenticator) GenerateToken(userID int, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.tokenTTL * time.Second).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(a.secretKey)
	if err != nil {
		a.logger.LogError("failed to sign token", map[string]interface{}{"package": "auth", "method": "GenerateToken"})
		return "", err
	}
	return signedToken, nil
}

// ValidateToken validates a JWT token and returns the claims.
func (a *jwtAuthenticator) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				a.logger.LogError("malformed token", map[string]interface{}{"package": "auth", "method": "ValidateToken"})
				return nil, errors.New("malformed token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				a.logger.LogError("token has expired", map[string]interface{}{"package": "auth", "method": "ValidateToken"})
				return nil, errors.New("token has expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				a.logger.LogError("token not valid yet", map[string]interface{}{"package": "auth", "method": "ValidateToken"})
				return nil, errors.New("token not valid yet")
			} else {
				a.logger.LogError("token validation error", map[string]interface{}{"package": "auth", "method": "ValidateToken"})
				return nil, errors.New("token validation error")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		a.logger.LogInfo("user token validated successfully", map[string]interface{}{"package": "auth", "method": "ValidateToken", "user_id": claims.UserID})
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
