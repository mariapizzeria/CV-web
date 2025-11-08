package education

import (
	"net/http"

	"github.com/mariapizzeria/cv-web/backend/pkg/customErrors"
	"github.com/mariapizzeria/cv-web/backend/pkg/response"
)

type Handler struct {
	Repository EducationRepository
}

type HandlerDeps struct {
	Repository EducationRepository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		Repository: deps.Repository,
	}
	router.HandleFunc("GET /education", handler.getAll())
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := handler.Repository.GetAll()
		if err != nil {
			customErrors.GetContentError(w)
			return
		}
		response.JsonEncoder(w, content, http.StatusOK)
	}
}
