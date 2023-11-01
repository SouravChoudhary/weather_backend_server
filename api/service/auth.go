package service

import (
	"time"
	"weatherapp/apperror"
	"weatherapp/db/user"
	"weatherapp/pkg/auth"
	"weatherapp/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

// Auth defines the auth service methods for user authentication.
type Auth interface {
	RegisterUser(username, password string, dob time.Time) error
	LoginUser(username, password string) (*user.User, string, error)
}

type userAuth struct {
	store   user.Store
	jwtAuth auth.JWTAuthenticator
	logger  logging.Logger
}

// NewAuthService creates a new AuthService with the provided dependencies.
func NewAuthService(store user.Store, jwtAuth auth.JWTAuthenticator, logger logging.Logger) Auth {
	return &userAuth{
		store:   store,
		jwtAuth: jwtAuth,
		logger:  logger,
	}
}

// RegisterUser registers a new user.
func (s *userAuth) RegisterUser(username, password string, dob time.Time) error {
	existingUser, _ := s.store.GetByUsername(username)
	if existingUser != nil {
		return apperror.UserAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.ReturnServiceErr("AS03", err)
	}

	newUser := &user.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		DateOfBirth:  dob,
	}
	err = s.store.Create(newUser)
	if err != nil {
		return err
	}
	s.logger.LogInfo("user registered successfully", map[string]interface{}{"user_id": newUser.ID})
	return nil
}

// LoginUser logs in a user and returns a JWT token.
func (s *userAuth) LoginUser(username, password string) (*user.User, string, error) {

	// Retrieve the user by username.
	user, err := s.store.GetByUsername(username)
	if err != nil {
		return nil, "", apperror.ReturnServiceErr("U01", err)
	}

	// Verify the password.
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", apperror.ReturnServiceErr("AS01", err)
	}

	// Generate a JWT token.
	token, err := s.jwtAuth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, "", apperror.ReturnServiceErr("AS02", err)
	}
	s.logger.LogInfo("user logged in successfully", map[string]interface{}{"user_id": user.ID})
	return user, token, nil
}
