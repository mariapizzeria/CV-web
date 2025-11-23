package experience

import (
	"cv-web/middleware"
	"cv-web/pkg/customErrors"
	"cv-web/pkg/response"
	"net/http"
	"strconv"

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
	router.Handle("POST /experience", middleware.IsAllowed(handler.createNewContent()))
	router.Handle("PUT /experience/{id}", middleware.IsAllowed(handler.updateContentById()))
	router.Handle("DELETE /experience/{id}", middleware.IsAllowed(handler.deleteContentById()))
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
