package models

import (
	"errors"
	"regexp"
	"time"
)

type RequestValidate interface {
	Validate() error
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (lr *LoginRequest) Validate() error {
	if lr.Username == "" || lr.Password == "" {
		return errors.New("username or password cannot be empty")
	}
	return nil
}

type RegistrationRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

func (rr *RegistrationRequest) Validate() error {

	if rr.Username != "" && !isValidEmail(rr.Username) {
		return errors.New("username must be a valid email address")
	}

	if rr.Password == "" {
		return errors.New("password cannot be empty")
	}

	if rr.DateOfBirth == "" {
		return errors.New("dob cannot be empty")
	}

	_, err := time.Parse("2006-01-02", rr.DateOfBirth)
	if err != nil {
		return errors.New("DateOfBirth must be in the format '2006-01-02'")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailPattern, email)
	return matched
}
