package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	// Mount your routes here, e.g.:
	// mux.HandleFunc("/api/v1/resource", app.handleResource)

	// posts

	// Comments

	// Users

	// Auth

	return mux
}

func (app *application) run() error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
