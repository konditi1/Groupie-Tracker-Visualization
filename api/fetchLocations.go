package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie/models"
)

func FetchLocations() (models.IndexLocations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error Fetching Locations Data", err)
		return models.IndexLocations{}, err
	}
	defer res.Body.Close()

	var locations models.IndexLocations
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		// fmt.Println("Error Decoding Locations Data", err)
		return models.IndexLocations{}, err
	}
	return locations, nil
}
