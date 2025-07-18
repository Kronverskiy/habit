package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kronverskiy/habit/api/gen"
	"github.com/Kronverskiy/habit/internal/rest"
	"github.com/Kronverskiy/habit/internal/rest/middleware"
	"github.com/go-chi/chi/v5"
)

func Run() {
	server := rest.NewHabitHandler()

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("load app config: %s", err.Error())
	}

	r := chi.NewMux()

	r.Use(middleware.Cors)

	h := gen.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
	}

	log.Fatal(s.ListenAndServe())
}
