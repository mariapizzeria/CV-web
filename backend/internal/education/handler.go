package education

import (
	"cv-web/backend/middleware"
	"cv-web/backend/pkg/customErrors"
	"cv-web/backend/pkg/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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
	router.Handle("POST /education", middleware.IsAllowed(handler.createNewContent()))
	router.Handle("PUT /education/{id}", middleware.IsAllowed(handler.updateContent()))
	router.Handle("DELETE /education/{id}", middleware.IsAllowed(handler.deleteContentById()))
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

func (handler *Handler) updateContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[Education](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		id := r.PathValue("id")
		uintId, err := strconv.ParseUint(id, 10, 32)

		res, err := handler.Repository.UpdateContent(&Education{
			gorm.Model{ID: uint(uintId)},
			body.Duration,
			body.College,
			body.Course,
		})
		if err != nil {
			customErrors.UpdateError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) deleteContentById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		uintId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		err = handler.Repository.DeleteContent(uint(uintId))
		if err != nil {
			return
		}
		response.JsonEncoder(w, err, http.StatusNoContent)
	}
}
