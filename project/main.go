package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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
	
	
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	
	apiCfg := apiConfig{
		DB: dbQueries,
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
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
	
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
	
	router.Mount("/v1", v1Router)
	
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	
	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)
	
	log.Printf("Serving on port: %s\n", portString)
	log.Fatal(srv.ListenAndServe())
}