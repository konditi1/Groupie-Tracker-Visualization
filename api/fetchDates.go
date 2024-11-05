package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie/models"
)

func FetchDates() (models.IndexDates, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error fetching Dates", err)
		return models.IndexDates{}, err
	}
	defer res.Body.Close()

	var dates models.IndexDates
	err = json.NewDecoder(res.Body).Decode(&dates)
	if err != nil {
		// fmt.Println("Error Decoding Dates", err)
		return models.IndexDates{}, err
	}
	return dates, nil
}
