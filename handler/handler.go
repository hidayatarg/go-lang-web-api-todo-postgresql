package handler

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

// Constructor
// Create new server and return a pointer to that server
func NewServer() *Server {
	return &Server{}
}

func SetupRouter() *chi.Mux {
	server := NewServer()

	// Get HTTP Req
	r := chi.NewRouter()

	setupMiddlewares(r)

	server.setupEndpoints(r)

	return r
}
