package blog

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/blog", h.handleAllBlogs).Methods("GET")
	router.HandleFunc("/blog/{id}", h.handleOneBlog).Methods("GET")
}

func (h *Handler) handleAllBlogs(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleOneBlog(w http.ResponseWriter, r *http.Request) {}
