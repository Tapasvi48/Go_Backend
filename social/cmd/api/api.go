package main

import (
	"net/http"
	"time"
)

type application struct {
	config config
}
type config struct {
	port string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	//
	mux.HandleFunc("/users", app.healthCheckHandler)
	return mux
}
func (app *application) startServer(mux *http.ServeMux) error {

	server := http.Server{
		Addr:         app.config.port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	return server.ListenAndServe()

}
