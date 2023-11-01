package apperror

import (
	"net/http"
)

type Error struct {
	Code        string
	Description string
	Details     map[string]interface{}
}

// Implement the error interface for MyError
func (e Error) Error() string {
	return e.Description
}

// StatusCode returns the HTTP status code for the error code
func (e Error) StatusCode() int {
	if status, ok := ErrorHTTPStatusMapping[e.Code]; ok {
		return status
	}
	// Return a default status code (e.g., 500 Internal Server Error) if the code is not found in the mapping.
	return http.StatusInternalServerError
}

var (
	ErrorCodeNotImplemented = Error{Code: "000", Description: "Error code not implemented", Details: nil}

	// User related errors [RANGE: U01-U10]
	UserNotFound        = Error{Code: "U01", Description: "User not found", Details: nil}
	UserCreationFailed  = Error{Code: "U02", Description: "User creation failed", Details: nil}
	UserDeletionFailed  = Error{Code: "U03", Description: "User deletion failed", Details: nil}
	UserRetrievalFailed = Error{Code: "U04", Description: "User retrieval failed", Details: nil}
	UserUpdationFailed  = Error{Code: "U05", Description: "User updation failed", Details: nil}

	// User Service errors [RANGE: US01-US10]
	UserAlreadyExist = Error{Code: "US01", Description: "User already in use", Details: nil}

	// Auth Service error [RANGE: A01-A10]
	InvalidPassword     = Error{Code: "AS01", Description: "Invalid password", Details: nil}
	GenerateTokenFail   = Error{Code: "AS02", Description: "Generate Token Fail", Details: nil}
	PasswordHashingFail = Error{Code: "AS03", Description: "Password hashing fail", Details: nil}

	// weather history record related errors [RANGE: WH01-WH10]
	WeatherHistoryRecordNotFound          = Error{Code: "WH01", Description: "weather history record not found", Details: nil}
	InvalidWeatherHistoryRecordID         = Error{Code: "WH02", Description: "invalid weather history record ID", Details: nil}
	FailedToCreateWeatherHistoryRecord    = Error{Code: "WH03", Description: "failed to create weather history record", Details: nil}
	FailedToDeleteWeatherHistoryRecord    = Error{Code: "WH04", Description: "failed to delete weather history record", Details: nil}
	PartialBulkDeleteWeatherHistoryRecord = Error{Code: "WH05", Description: "partial bulk delete weather history record", Details: nil}

	// weather api client error
	WeatherApiClientfail = Error{Code: "WC01", Description: "weather api client call failure", Details: nil}
)
