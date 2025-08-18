package main

import (
	"log"
	"net/http"
)

func (app *App) RunServer() {
	srv := &http.Server{
		Addr:    app.Addr,
		Handler: app.Routes(),
	}

	log.Printf("Starting server on %s", app.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
