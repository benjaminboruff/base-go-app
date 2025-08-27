package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.SessionManager.Put(r.Context(), "message", "Hello from a session!")

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message, use
	// the http.Error() function to send an Internal Server Error response to the
	// user, and then return from the handler so no subsequent code is executed.
	// ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		app.HTMLDir + "/base.tmpl",
		app.HTMLDir + "/partials/nav.tmpl",
		app.HTMLDir + "/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ProfileView handler
func (app *App) ProfileView(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")
	msg := app.SessionManager.GetString(r.Context(), "message")
	log.Println(msg)
	files := []string{
		app.HTMLDir + "/base.tmpl",
		app.HTMLDir + "/partials/nav.tmpl",
		app.HTMLDir + "/pages/user.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	user, err := app.Env.users.Show(id)
	if err != nil || id < 1 {
		log.Print(err.Error())
		http.NotFound(w, r)
		return
	}

	err = ts.ExecuteTemplate(w, "base", user)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ProfileCreate handler
func (app *App) ProfileCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new profile ..."))
}

// ProfileCreatePost handler
func (app *App) ProfileCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new profile ..."))
}
