package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"groupie/models"
)

// Mock FetchLocations function to avoid calling the actual API
func init() {
	// Mocking FetchLocations to avoid external calls
	Locations = models.IndexLocations{
		Index: []models.Locations{
			{Id: 1, Locations: []string{"Location 1", "Location 2"}, Dates: "2024-09-15"},
			{Id: 2, Locations: []string{"Location 3", "Location 4"}, Dates: "2024-09-15"},
		},
	}
	fetchLocationErr = nil

	// Mocking FetchArtists
	Artist = []models.Artist{
		{Name: "Queen", Id: 1},
		{Name: "The Beatles", Id: 2},
	}
}

// TestLocationHandler tests the LocationHandler function
func TestLocationHandler(t *testing.T) {
	// Define test cases
	tests := []struct {
		name               string
		method             string
		query              string
		expectedStatusCode int
	}{
		{
			name:               "Invalid artist ID",
			method:             "GET",
			query:              "artistId=999&name=Queen&image=someImageURL",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Invalid artist name",
			method:             "GET",
			query:              "artistId=1&name=InvalidArtist&image=someImageURL",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Method not allowed",
			method:             "POST",
			query:              "artistId=1&name=Queen&image=someImageURL",
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}

	// Loop over the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new HTTP request with the provided query parameters
			req, err := http.NewRequest(tc.method, "/location?"+tc.query, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// Create a new ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the LocationHandler function
			handler := http.HandlerFunc(LocationHandler)
			handler.ServeHTTP(rr, req)

			// Check if the status code matches the expected status code
			if status := rr.Code; status != tc.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatusCode)
			}
		})
	}
}
