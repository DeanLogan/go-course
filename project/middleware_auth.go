package main

import (
	"fmt"
	"net/http"
	
	"github.com/DeanLogan/go-course/project/internal/auth"
	"github.com/DeanLogan/go-course/project/internal/database"
)


type authedHandler func (http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Auth error: %v", err))
			return 
		}
		
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, 400, fmt.Sprintf("Could not get user: %v", err))
			return 
		}
		
		handler(w, r, user)
	}
}