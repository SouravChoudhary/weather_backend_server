package apperror

import "net/http"

var ErrorMap = map[string]Error{
	"U01": UserNotFound,
	"U02": UserCreationFailed,
	"U03": UserDeletionFailed,
	"U04": UserRetrievalFailed,
	"U05": UserUpdationFailed,

	"US01": UserAlreadyExist,

	"AS01": InvalidPassword,
	"AS02": GenerateTokenFail,
	"AS03": PasswordHashingFail,

	"WH01": WeatherHistoryRecordNotFound,
	"WH02": InvalidWeatherHistoryRecordID,
	"WH03": FailedToCreateWeatherHistoryRecord,
	"WH04": FailedToDeleteWeatherHistoryRecord,
	"WH05": PartialBulkDeleteWeatherHistoryRecord,

	"WC01": WeatherApiClientfail,
}

var ErrorHTTPStatusMapping = map[string]int{ // we can have http status directly inside MyError struct.
	"U01": http.StatusNotFound,            // 404 Not Found
	"U02": http.StatusInternalServerError, // 500 Internal Server Error
	"U03": http.StatusInternalServerError, // 500 Internal Server Error
	"U04": http.StatusInternalServerError, // 500 Internal Server Error
	"U05": http.StatusInternalServerError, // 500 Internal Server Error

	"US01": http.StatusConflict, // 409 Conflict due to the existing resource

	"AS01": http.StatusBadRequest,          // 400 Bad Request
	"AS02": http.StatusInternalServerError, // 500 Internal Server Error
	"AS03": http.StatusInternalServerError,

	"WH01": http.StatusNotFound,            // 404 Not Found
	"WH02": http.StatusBadRequest,          // 400 Bad Request
	"WH03": http.StatusInternalServerError, // 500 Internal Server Error
	"WH04": http.StatusInternalServerError, // 500 Internal Server Error
	"WH05": http.StatusPartialContent,      // 206 Partial Delete operation Error

	"WC01": http.StatusInternalServerError,
}
