package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := OpenDatabase()
	if err != nil {
		log.Printf("error opening database connection %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello New World!"))
	})

	r.Get("/Items/{ItemName}", get)

	http.ListenAndServe("localhost:8080", r)
}
