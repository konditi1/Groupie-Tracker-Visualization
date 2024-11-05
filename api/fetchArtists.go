package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie/models"
)


func FetchArtists() ([]models.Artist, error) {
	var res, FetchArtistsErr = http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if FetchArtistsErr != nil {
		fmt.Println("Error Fetching Artist", FetchArtistsErr)
		return nil, FetchArtistsErr
	}
	defer res.Body.Close()

	var artists []models.Artist
	err := json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Error Decording Artist", err)
		return nil, err
	}
	return artists, nil
}
