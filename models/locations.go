package models

type IndexLocations struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id        int    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string `json:"dates"`
}

type LocationWithArtist struct {
    Id        int
    Locations []string
    Dates     string
    Name      string
    Image     string
}
