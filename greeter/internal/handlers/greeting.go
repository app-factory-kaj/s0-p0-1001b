package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"greeter/internal/gen"
)

const defaultName = "World"

// Server implements gen.ServerInterface.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetGreeting(w http.ResponseWriter, r *http.Request, params gen.GetGreetingParams) {
	name := params.Name
	if name == "" {
		name = defaultName
	}

	greeting := gen.Greeting{
		Name:    name,
		Message: fmt.Sprintf("Hello, %s!", name),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(greeting)
}
