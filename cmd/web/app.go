package main

// App defines a struct to hold applications-wide dependencies and
// application settings.
type App struct {
	Addr    string
	HTMLDir string
	DistDir string
}
