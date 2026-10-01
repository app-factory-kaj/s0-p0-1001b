package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	srv := handlers.NewServer()
	gen.HandlerFromMux(srv, r)

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", r); err != nil {
		log.Fatal(err)
	}
}
