package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie/api"
	"groupie/models"
)

var Locations, fetchLocationErr = api.FetchLocations()

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	artistIdStr := r.URL.Query().Get("artistId")
	artistName := r.URL.Query().Get("name")
	artistImage := r.URL.Query().Get("image")


	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	if artistId < 1 || artistId > len(Artist) {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist Id provided")
		fmt.Println("Invalid artist Id provided")
		return
	}

	check := false
	id := 1
	for i, name := range Artist {
		if name.Name == artistName {
			id = i + 1
			check = true
			break
		}
	}
	if id != artistId {
		ErrorPage(w, r, http.StatusBadRequest, "Mismatch artist Id and artist Name provided")
		fmt.Println("Mismatch artist Id and artist Name provided")
		return
	}

	if !check {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist name")
		fmt.Println("Invalid artist name")
		return
	}

	location := Locations.Index[artistId-1]
	if fetchLocationErr != nil {
		fmt.Println("Error fetching the locations Id", fetchLocationErr)
		ErrorPage(w, r, http.StatusBadRequest, "Error fetching locations Id")
		return
	}

	LocationsData := struct {
		Name      string
		Image     string
		Id        int
		Locations models.Locations
	}{
		Name:      artistName,
		Image:     artistImage,
		Id:        artistId,
		Locations: location,
	}

	execErr := tmpl.ExecuteTemplate(w, "locations.html", LocationsData)
	if execErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering locations template")
		return
	}
}
