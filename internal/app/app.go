package app

import (
	"log"
	"net/http"

	"github.com/Kronverskiy/habit/api/gen"
	"github.com/Kronverskiy/habit/internal/rest"
	"github.com/Kronverskiy/habit/internal/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func Run() {
	server := rest.NewHabitHandler()

	r := chi.NewMux()

	r.Use(middleware.Cors)

	h := gen.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
