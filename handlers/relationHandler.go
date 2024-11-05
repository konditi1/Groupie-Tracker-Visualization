package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie/api"
	"groupie/models"
)

var Relations, fetchRelationErr = api.FetchRelation()

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	artistName := r.URL.Query().Get("name")
	artistIdStr := r.URL.Query().Get("artistId")

	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	if artistId < 1 || artistId > len(Relations.Index) {
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

	if fetchRelationErr != nil {
		ErrorPage(w, r, http.StatusBadRequest, "Error fetching relations data")
		return
	}

	relation := Relations.Index[artistId-1]
	relationsData := struct {
		Name     string
		Relation models.Relation
	}{
		Name:     artistName,
		Relation: relation,
	}
	execErr := tmpl.ExecuteTemplate(w, "relation.html", relationsData)
	if execErr != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "Error rendering realations template")
	}
}
