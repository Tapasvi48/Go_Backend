package main

import "log"

func main() {

	app := &application{
		config: config{
			port: ":8080",
		},
	}
	mux := app.mount()
	err := app.startServer(mux)
	if err != nil {
		log.Fatal("Server can't start error")
		return
	}
	log.Printf("%s server started", app.config.port)

}
