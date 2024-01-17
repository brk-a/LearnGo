package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/brk-a/0x05-rss-feeds/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	AD *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString=="" {
		log.Fatal("port could not be found in env")
	}

	dbUrlString := os.Getenv("DB_URL")
	if dbUrlString=="" {
		log.Fatal("db URL could not be found in env")
	}

	conn, err := sql.Open("postgres", dbUrlString)
	if err!=nil {
		log.Fatal("could not connect to database", err)
	}

	apiCfg := apiConfig{
		DB: database.new(conn),
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
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	log.Printf("server starting at port %v", portString)
	err := srv.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}
