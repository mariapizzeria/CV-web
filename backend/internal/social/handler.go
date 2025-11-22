package social

import (
	"cv-web/backend/middleware"
	"cv-web/backend/pkg/customErrors"
	"cv-web/backend/pkg/response"
	"net/http"
	"strconv"

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
		Repository: deps.Repository,
	}
	router.HandleFunc("GET /social", handler.getAll())
	router.Handle("POST /social", middleware.IsAllowed(handler.createNew()))
	router.Handle("PUT /social/{id}", middleware.IsAllowed(handler.updateContent()))
	router.Handle("DELETE /social/{id}", middleware.IsAllowed(handler.deleteContent()))
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

func (handler *Handler) createNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[SocialResponse](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		res, err := handler.Repository.CreateNewContent(&SocialResponse{
			body.Link,
			body.Image,
		})
		if err != nil {
			customErrors.GetContentError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusOK)
	}
}

func (handler *Handler) updateContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := response.HandleBody[Social](w, r)
		if err != nil {
			customErrors.ReadBodyError(w)
			return
		}
		id := r.PathValue("id")
		uintId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		res, err := handler.Repository.UpdateContent(&Social{
			Model: gorm.Model{ID: uint(uintId)},
			Link:  body.Link,
			Image: body.Image,
		})
		if err != nil {
			customErrors.UpdateError(w)
			return
		}
		response.JsonEncoder(w, res, http.StatusAccepted)
	}
}

func (handler *Handler) deleteContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		uintId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			customErrors.ServerError(w)
			return
		}
		err = handler.Repository.DeleteContent(uint(uintId))
		response.JsonEncoder(w, err, http.StatusAccepted)
	}
}
