package server

import (
	"github.com/AHMED-D007A/Blogging-Platform-API/internal/service/blog"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", blog.HandleAllBlogs).Methods("GET")
	router.HandleFunc("/posts/{id}", blog.HandleOneBlog).Methods("GET")
}
