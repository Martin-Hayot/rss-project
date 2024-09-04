package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Martin-Hayot/rss-project/internal/database"
	"github.com/google/uuid"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]bool{"healthy": true})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), params.ApiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
