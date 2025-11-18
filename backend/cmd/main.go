package main

import (
	"log"
	"net/http"

	"github.com/mariapizzeria/cv-web/backend/configs"
	"github.com/mariapizzeria/cv-web/backend/db"
	"github.com/mariapizzeria/cv-web/backend/internal/education"
	"github.com/mariapizzeria/cv-web/backend/internal/experience"
	"github.com/mariapizzeria/cv-web/backend/internal/social"
	"github.com/mariapizzeria/cv-web/backend/internal/stack"
	"github.com/mariapizzeria/cv-web/backend/middleware"
)

func main() {
	conf := configs.LoadConfig()
	newDB := db.NewDb(conf)
	router := http.NewServeMux()

	// repository
	educationRepository := education.NewRepository(newDB)
	experienceRepository := experience.NewRepository(newDB)
	socialRepository := social.NewRepository(newDB)
	skillsRepository := stack.NewRepository(newDB)
	// handlers
	education.NewHandler(router, education.HandlerDeps{
		Repository: educationRepository,
	})
	experience.NewHandler(router, experience.HandlerDeps{
		Repository: experienceRepository,
	})
	social.NewHandler(router, social.HandlerDeps{
		Repository: socialRepository,
	})
	stack.NewHandler(router, stack.HandlerDeps{
		Repository: skillsRepository,
	})
	// server
	server := &http.Server{
		Addr:    ":8082",
		Handler: middleware.CORSHandler(router),
	}
	log.Println("Listening on port 8082")
	server.ListenAndServe()
}
