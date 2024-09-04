package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Martin-Hayot/rss-project/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set and/or is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL must be set and/or is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthy", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Get("/users", apiCfg.handlerGetUser)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on port %v", port)

	error := srv.ListenAndServe()

	if error != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", port)
}
