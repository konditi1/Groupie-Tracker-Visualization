package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"groupie/models"
)

func ArtistFullProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist Id")
		fmt.Println("Invalid artist Id")
		return
	}

	var artistFullProfile string

	// Parse the query parameters
	queryParams, queryParamErr := url.ParseQuery(r.URL.RawQuery)
	if queryParamErr != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid URL")
		return
	}

	artistName := queryParams.Get("name")
	artistImage := queryParams.Get("image")
	artistIdStr := queryParams.Get("artistId")

	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist Id")
		fmt.Println("Invalid artist Id")
		return
	}

	if artistId < 1 || artistId > len(Artist) {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist Id provided")
		fmt.Println("Invalid artist Id provided")
		return
	}

	Id := 1
	check := false
	for i, name := range Artist {
		if name.Name == artistName {
			Id = i + 1
			check = true
			break
		}
	}

	if Id != artistId {
		ErrorPage(w, r, http.StatusBadRequest, "Mismatch artist Id and artist Name provided")
		fmt.Println("Mismatch artist Id and artist Name provided")
		return
	}

	if !check {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist name")
		fmt.Println("Invalid artist namee")
		return
	}

	// Validate the image URL format
	imageURLPattern := `^https://groupietrackers\.herokuapp\.com/api/images/[\w-]+\.jpeg$`
	matched, _ := regexp.MatchString(imageURLPattern, artistImage)
	if !matched {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid image URL format")
		fmt.Println("Invalid image URL format")
		return
	}

	res1, _ := http.Get(artistImage)
	if res1.StatusCode != 200 {
		ErrorPage(w, r, http.StatusInternalServerError, "Image Url tampered with")
		fmt.Println("Image url tampered with")
		return
	}

	date := Dates.Index[artistId-1]
	location := Locations.Index[artistId-1]
	artist := Artist[artistId-1]
	relation := Relations.Index[artistId-1]

	artistFullProfileInfo := struct {
		Title     string
		Artist    models.Artist
		Name      string
		Image     string
		Dates     models.Dates
		Locations models.Locations
		Relation  models.Relation
	}{
		Title:     artistFullProfile,
		Artist:    artist,
		Name:      artistName,
		Image:     artistImage,
		Dates:     date,
		Locations: location,
		Relation:  relation,
	}
	execErr := tmpl.ExecuteTemplate(w, "artistFullProfile.html", artistFullProfileInfo)
	if execErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering artistFullProfile template")
		fmt.Println("Error rendering artistFullProfile template")
	}
}
