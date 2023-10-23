package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/machingclee/rssagg/internal/auth"
	"github.com/machingclee/rssagg/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
	}
	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}
	responseWithJson(w, 200, databaseUsertoUser(&user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	responseWithJson(w, 200, databaseUsertoUser(&user))
}
