package models

type IndexRelation struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	Id             int               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
