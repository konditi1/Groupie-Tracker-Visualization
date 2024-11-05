package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"groupie/api"
	"groupie/models"
)

type Suggestion struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the query parameter from the URL
	query := r.URL.Query().Get("query")
	if query == "" {
		ErrorPage(w, r, http.StatusBadRequest, "Query parameter is required")
		return
	}

	// Fetch artists from the API
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artist name")
		ErrorPage(w, r, http.StatusInternalServerError, "Error fetching artists name")
		return
	}

	// Filter artists based on the query
	var artistsFiltered []models.Artist
	var membersFiltered []models.Artist
	var creationDatesFiltered []models.Artist
	var firstAlbumsFiltered []models.Artist
	var suggestions []Suggestion

	// Convert the query to lowercase once for efficiency
	queryLower := strings.ToLower(query)

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), queryLower) {
			artistsFiltered = append(artistsFiltered, artist)
			suggestions = append(suggestions, Suggestion{Name: artist.Name, Type: "artist/band"})
		} else if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), queryLower) {
			creationDatesFiltered = append(creationDatesFiltered, artist)
			suggestions = append(suggestions, Suggestion{Name: artist.Name, Type: "creation date"})
		} else if strings.Contains(strings.ToLower(artist.FirstAlbum), queryLower) {
			firstAlbumsFiltered = append(firstAlbumsFiltered, artist)
			suggestions = append(suggestions, Suggestion{Name: artist.Name, Type: "first album"})
		}

		for i := 0; i < len(artist.Members); i++ {
			if strings.Contains(strings.ToLower(artist.Members[i]), queryLower) {
				membersFiltered = append(membersFiltered, artist)
				suggestions = append(suggestions, Suggestion{Name: artist.Members[i], Type: "member"})
				break
			}
		}
	}

	// Fetch locations from the API
	locations, err := api.FetchLocations()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error fetching locations")
		return
	}

	// Prepare a map for quick lookup of artists by Id
	artistMap := make(map[int]models.Artist)
	for _, artist := range artists {
		artistMap[artist.Id] = artist
	}

	var locationsFiltered []models.LocationWithArtist

	// Filter locations based on the query
	for _, location := range locations.Index {
		for _, loc := range location.Locations {
			if strings.Contains(strings.ToLower(loc), queryLower) {
				// Check if there's an artist matching the location Id
				if artist, found := artistMap[location.Id]; found {
					locationsFiltered = append(locationsFiltered, models.LocationWithArtist{
						Id:        location.Id,
						Locations: location.Locations,
						Dates:     location.Dates,
						Name:      artist.Name,
						Image:     artist.Image,
					})
				}
				suggestions = append(suggestions, Suggestion{Name: loc, Type: "location"})
				break
			}
		}
	}

	// execute a JSON request (for dynamic suggestions), return JSON response
	if r.Header.Get("Accept") == "application/json" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(suggestions)
        return
    }

	searchData := struct {
		Artists       []models.Artist
		Members       []models.Artist
		CreationDates []models.Artist
		FirstAlbums   []models.Artist
		Locations     []models.LocationWithArtist
	}{
		Artists:       artistsFiltered,
		Members:       membersFiltered,
		CreationDates: creationDatesFiltered,
		FirstAlbums:   firstAlbumsFiltered,
		Locations:     locationsFiltered,
	}

	// Execute the full search result template
	if err = tmpl.ExecuteTemplate(w, "search.html", searchData); err != nil {
		fmt.Println("Template execution error:", err)
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering search artistsFiltered")
		return
	}
}
