package handler

import "github.com/go-chi/chi"

type Server struct {
}

func (s *Server) setupEndpoints(r *chi.Mux) {

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

	server.setupEndoints(r)

	return r
}
