package main

import (
	"net/http"
)

func (app *App) Routes() http.Handler {

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir(app.DistDir))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("GET /dist/", http.StripPrefix("/dist", fileServer))

	// Register other application routes as normal.
	mux.HandleFunc("GET /{$}", app.Home) // Restrict route to exact match on /
	mux.HandleFunc("GET /profile/view/{id}", app.ProfileView)
	mux.HandleFunc("GET /profile/create", app.ProfileCreate)
	mux.HandleFunc("POST /profile/create", app.ProfileCreatePost)

	// Application routes for User
	mux.HandleFunc("GET /user/signup", app.SignupUser)
	mux.HandleFunc("GET /user/login", app.LoginUser)
	mux.HandleFunc("POST /user/login", app.VerifyUser)

	return app.SessionManager.LoadAndSave(mux)
}
