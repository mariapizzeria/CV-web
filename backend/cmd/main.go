package main

import (
	"log"
	"net/http"

	"github.com/mariapizzeria/cv-web/backend/configs"
	"github.com/mariapizzeria/cv-web/backend/db"
	"github.com/mariapizzeria/cv-web/backend/internal/education"
)

func main() {
	conf := configs.LoadConfig()
	newDB := db.NewDb(conf)
	router := http.NewServeMux()
	router.HandleFunc("GET /stack", hello)

	// repository
	educationRepository := education.NewRepository(newDB)
	// handlers
	education.NewHandler(router, education.HandlerDeps{
		Repository: educationRepository,
	})
	// server
	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	log.Println("Listening on port 8082")
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Print("hello")
}
