package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie/api"
	"groupie/models"
)

var Dates, fetchDatesErr = api.FetchDates()

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")

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

	artistName := r.URL.Query().Get("name")
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

	if fetchDatesErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error fetching dates data")
		return
	}

	date := Dates.Index[artistId-1]

	DatesData := struct {
		Dates models.Dates
		Name  string
	}{
		Dates: date,
		Name:  artistName,
	}

	execErr := tmpl.ExecuteTemplate(w, "dates.html", DatesData)
	if execErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering dates template")
	}
}
