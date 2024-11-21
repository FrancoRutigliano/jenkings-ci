package main

import (
	"jenkings-ci/config"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	config.Setup()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola crayola"))
	})

	log.Println(os.Getenv("DATABASE_URL"))
	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", r)
}
