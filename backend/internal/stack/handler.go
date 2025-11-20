package stack

import (
	"net/http"
	"strconv"

	"github.com/mariapizzeria/cv-web/backend/middleware"
	"github.com/mariapizzeria/cv-web/backend/pkg/customErrors"
	"github.com/mariapizzeria/cv-web/backend/pkg/response"
	"gorm.io/gorm"
)

type Handler struct {
	Repository *Repository
}

type HandlerDeps struct {
	Repository *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{
		deps.Repository,
	}
	router.HandleFunc("GET /skills", handler.getAll())
	router.Handle("POST /skills", middleware.IsAllowed(handler.createContent()))
	router.Handle("PUT /skills/{id}", middleware.IsAllowed(handler.updateContent()))
	router.Handle("DELETE /skills/{id}", middleware.IsAllowed(handler.deleteContent()))
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

func (handler *Handler) createContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[StackResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		res, err := handler.Repository.CreateNewContent(&StackResponse{
			Type:  body.Type,
			Skill: body.Skill,
		})
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) updateContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[StackResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		res, err := handler.Repository.UpdateContent(&Stack{
			Model: gorm.Model{ID: uint(idUint)},
			Type:  body.Type,
			Skill: body.Skill,
		})
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) deleteContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		err = handler.Repository.DeleteContent(uint(idUint))
		response.JsonEncoder(w, err, http.StatusAccepted)
	}
}
