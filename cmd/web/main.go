package main

import ()

func main() {

	app := &App{
		Addr:    ":8080",
		HTMLDir: "./ui/html",
		DistDir: "./ui/dist",
	}

	app.RunServer()
}
