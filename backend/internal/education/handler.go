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
	router.HandleFunc("POST /education", handler.createNewContent())
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

func (handler *Handler) createNewContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[EducationResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		res, err := handler.Repository.CreateNewContent(&EducationResponse{
			body.Duration,
			body.College,
			body.Course,
		})
		if err != nil {
			customErrors.CreateError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusOK)
	}
}
