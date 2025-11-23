package main

import (
	"cv-web/internal/education"
	"cv-web/internal/experience"
	"cv-web/internal/social"
	"cv-web/internal/stack"
	"cv-web/middleware"
	"cv-web/pkg/configs"
	"cv-web/pkg/db"
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
