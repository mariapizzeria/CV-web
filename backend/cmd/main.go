package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

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

	// static execute
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := strings.ToLower(filepath.Ext(r.URL.Path))
		switch ext {
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".svg":
			w.Header().Set("Content-Type", "image/svg+xml")
		case ".ico":
			w.Header().Set("Content-Type", "image/x-icon")
		}
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./static/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	}))
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
