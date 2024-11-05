package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie/models"
)

func FetchRelation() (models.IndexRelation, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error Fetching Relation Data", err)
		return models.IndexRelation{}, err
	}
	defer res.Body.Close()

	var relation models.IndexRelation
	err = json.NewDecoder(res.Body).Decode(&relation)
	if err != nil {
		// fmt.Println("Error Decoding Relation Data", err)
		return models.IndexRelation{}, err
	}
	return relation, nil
}
