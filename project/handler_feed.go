package main

import (
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/DeanLogan/go-course/project/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	
	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:	   params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
