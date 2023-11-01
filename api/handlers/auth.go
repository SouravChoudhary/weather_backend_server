package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"weatherapp/api/service"
	"weatherapp/models"
	"weatherapp/pkg/logging"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authService service.Auth
	logger      logging.Logger
}

func NewAuthHandler(authService service.Auth, logger logging.Logger) AuthHandler {
	return &authHandler{authService, logger}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid request data:" + err.Error()})
		return
	}

	if err := loginRequest.Validate(); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Validation error: " + err.Error()})
		return
	}

	user, token, err := h.authService.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, models.Response{Error: "Login failed"})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: map[string]interface{}{"Authorization": "Bearer " + token, "user_id": user.ID}})
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registrationRequest models.RegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&registrationRequest); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid request data"})
		return
	}

	if err := registrationRequest.Validate(); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Validation error: " + err.Error()})
		return
	}

	dateOfBirth, err := time.Parse("2006-01-02", registrationRequest.DateOfBirth)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid date_of_birth format"})
		return
	}

	err = h.authService.RegisterUser(registrationRequest.Username, registrationRequest.Password, dateOfBirth)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: "Registration failed"})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: "Registration successful"})
}
