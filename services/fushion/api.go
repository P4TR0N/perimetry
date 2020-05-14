package fushion

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
)


const host = "http://localhost:8080"

var apiKey = ""
var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

var baseURL, _ = url.Parse(host)

// Construct a new FusionAuth Client
var auth = fusionauth.NewClient(httpClient, baseURL, apiKey)

func FushionLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials fusionauth.LoginRequest
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&credentials)
	// Use FusionAuth Go client to log in the user

	authResponse, errors, err := auth.Login(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Write the response from the FusionAuth client as JSON
	var responseJSON []byte
	if errors != nil {
		responseJSON, err = json.Marshal(errors)
	} else {
		responseJSON, err = json.Marshal(authResponse)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func FushionRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials fusionauth.RegistrationRequest
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&credentials)
	// Use FusionAuth Go client to log in the user

	registerResponse, errors, err := auth.Register("", credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Write the response from the FusionAuth client as JSON
	var responseJSON []byte
	if errors != nil {
		responseJSON, err = json.Marshal(errors)
	} else {
		responseJSON, err = json.Marshal(registerResponse)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}