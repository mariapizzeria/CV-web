package experience

import (
	"net/http"
	"strconv"

	"github.com/mariapizzeria/cv-web/backend/pkg/customErrors"
	"github.com/mariapizzeria/cv-web/backend/pkg/response"
	"gorm.io/gorm"
)

type Handler struct {
	Repository *ExpRepository
}

type HandlerDeps struct {
	Repository *ExpRepository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Repository: deps.Repository,
	}
	router.HandleFunc("GET /experience", handler.getAll())
	router.HandleFunc("POST /experience", handler.createNewContent())
	router.HandleFunc("PUT /experience/{id}", handler.updateContentById())
	router.HandleFunc("DELETE /experience/{id}", handler.deleteContentById())
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := handler.Repository.GetAll()
		if err != nil {
			customErrors.GetContentError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusOK)
	}
}

func (handler *Handler) createNewContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[ExperienceResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		res, err := handler.Repository.CreateNewContent(&ExperienceResponse{
			body.Duration,
			body.Title,
			body.Description,
		})
		if err != nil {
			customErrors.GetContentError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) updateContentById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[ExperienceResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		res, err := handler.Repository.UpdateContent(&Experience{
			gorm.Model{ID: uint(idUint)},
			body.Duration,
			body.Title,
			body.Description,
		})
		if err != nil {
			customErrors.GetContentError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) deleteContentById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			customErrors.NotFound(w)
			return
		}
		err = handler.Repository.DeleteContent(uint(idUint))
		response.JsonEncoder(w, err, http.StatusAccepted)
	}
}
