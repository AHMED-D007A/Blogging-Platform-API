package server

import (
	"github.com/AHMED-D007A/Blogging-Platform-API/internal/service/blog"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", blog.CreateBlog).Methods("POST")
	router.HandleFunc("/posts/{id}", blog.CreateBlog).Methods("PUT")
	router.HandleFunc("/posts/{id}", blog.CreateBlog).Methods("DELETE")
	router.HandleFunc("/posts", blog.GetAllBlogs).Methods("GET")
	router.HandleFunc("/posts/{id}", blog.GetBlog).Methods("GET")
}
