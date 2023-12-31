package main

import (
	"encoding/json"
	"net/http"
	_"strconv"
	"time"

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

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// limitStr := r.URL.Query().Get("limit")
	// limit := 10
	// if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
	// 	limit = specifiedLimit
	// }
	
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
		//Limit:  int32(limit),
	})
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}
	
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}