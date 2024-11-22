package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	DatabaseURL string
	OtherVar    string
)

func main() {

	log.Println("env database: ", DatabaseURL)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola crayola"))
	})
	log.Println("Server started on :3000 ....")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("error starting the server, error: ", err)
	}
}
