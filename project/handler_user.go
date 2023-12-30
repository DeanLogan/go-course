package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"github.com/DeanLogan/go-course/project/internal/auth"
	"github.com/DeanLogan/go-course/project/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	
	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithErr(w, 403, fmt.Sprintf("Auth error: %v", err))
		return 
	}
	
	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Could not get user: %v", err))
		return 
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}