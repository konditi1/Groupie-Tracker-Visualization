package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/api"
)

var (
	Artist, fetchArtistErr = api.FetchArtists()
)

var tmpl *template.Template

func init() {
	// Parse the error HTML template file
	var err error
	tmpl, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println("Error parsing template: " + err.Error())
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	// Validate URL path
	validPaths := map[string]bool{
		"/":                  true,
		"/dates":             true,
		"/locations":         true,
		"/relations":         true,
		"/artistFullProfile": true,
		"/favicon.ico":       true,
	}

	if !validPaths[r.URL.Path] {
		ErrorPage(w, r, http.StatusNotFound, "Page Not Found")
		fmt.Println(r.URL.Path)
		fmt.Println("Artist page Not Found")
		return
	}

	if r.Method != "GET" {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if fetchArtistErr != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Error fetching artist data")
		return
	}

	execErr := tmpl.ExecuteTemplate(w, "index.html", Artist)
	if execErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering index template")
	}
}
