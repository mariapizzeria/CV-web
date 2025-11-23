package main

import (
	"cv-web/backend/internal/education"
	"cv-web/backend/internal/experience"
	"cv-web/backend/internal/social"
	"cv-web/backend/internal/stack"
	"cv-web/backend/middleware"
	"cv-web/backend/pkg/configs"
	"cv-web/backend/pkg/db"
	"log"
	"net/http"
)

func App() http.Handler {
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
	return router
}

func main() {
	app := App()
	server := &http.Server{
		Addr:    ":8082",
		Handler: middleware.CORSHandler(app),
	}
	log.Println("Server is listening")
	server.ListenAndServe()
}
