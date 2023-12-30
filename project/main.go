package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DeanLogan/go-course/project/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){
	godotenv.Load()
	
	portString := os.Getenv("PORT") // gets value from env using the key
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	
	dbURL := os.Getenv("DB_URL") // gets value from env using the key
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	
	
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database",err)
	}
	
	queries := database.New(conn)
	
	apiCfg := apiConfig{
		DB: queries,
	}
	
	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // scopes the handler to only "fire" on GET requests
	v1Router.Get("/err",handleErr)
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	router.Mount("/v1", v1Router)
	
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Port:", portString)
}