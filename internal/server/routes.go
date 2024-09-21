package server

import (
	"database/sql"

	"github.com/AHMED-D007A/Blogging-Platform-API/internal/service/blog"
	"github.com/gorilla/mux"
)

func RegisterBlogRoutes(router *mux.Router, db *sql.DB) {
	handler := blog.NewBlogHandler(*blog.NewBlogStorage(db))

	router.HandleFunc("/posts", handler.CreateBlogHandler).Methods("POST")
	router.HandleFunc("/posts/{id}", handler.UpdateBlogHandler).Methods("PUT")
	router.HandleFunc("/posts/{id}", handler.DeleteBlogHandler).Methods("DELETE")
	router.HandleFunc("/posts", handler.GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", handler.GetBlogHandler).Methods("GET")
}
