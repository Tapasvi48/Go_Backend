package main

import (
	user "backend/users"
	"log"
	"net/http"
)

type app struct {
	port string
}

// method to start the app

func (s *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	app := &app{port: ":8080"}
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    app.port,
		Handler: mux,
		// IdleTimeout: time.Minute,
		// ReadTimeout: 10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}
	mux.HandleFunc("/users", user.HandleGetUsers)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
