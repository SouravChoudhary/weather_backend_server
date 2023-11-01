package service

import (
	"weatherapp/db/user"
)

// UserService defines the user service methods.
type UserService interface {
	GetUserProfile(userID int) (*user.User, error)
}

type userService struct {
	store user.Store
}

// NewUserService creates a new UserService with the provided dependencies.
func NewUserService(store user.Store) UserService {
	return &userService{
		store: store,
	}
}

// GetUserProfile retrieves the user profile by userID.
func (s *userService) GetUserProfile(userID int) (*user.User, error) {
	user, err := s.store.GetByID(userID)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return user, nil
}
