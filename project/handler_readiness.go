package main 

import (
	"net/http"
	"fmt"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is interesting")
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 400, "Something went wrong")
}