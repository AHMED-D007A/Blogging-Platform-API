package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AHMED-D007A/Blogging-Platform-API/service/blog"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1/").Subrouter()

	s.handleRoutes(subrouter)

	log.Printf("Server is up and running on port: %v", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func (s *APIServer) handleRoutes(router *mux.Router) {
	blogHandler := blog.NewHandler()
	blogHandler.RegisterRoutes(router)
}
