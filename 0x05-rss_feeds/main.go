package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString=="" {
		log.Fatal("port could not be found in env")
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
