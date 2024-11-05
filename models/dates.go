package models

type IndexDates struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	Id        int    `json:"id"`
	Dates []string `json:"dates"`
}
