package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie/models"
)

// Mock function to simulate FetchArtists
func mockFetchArtists() ([]models.Artist, error) {
	return []models.Artist{
		{Name: "Queen", Image: "http://example.com/queen.jpg"},
	}, nil
}

// A modified version of ArtistHandler to accept the fetch function as a parameter
func ArtistHandlerWithMock(fetchArtistsFunc func() ([]models.Artist, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate URL path
		validPaths := map[string]bool{
			"/":                  true,
			"/dates":             true,
			"/locations":         true,
			"/relations":         true,
			"/artistFullProfile": true,
			"/favicon.ico":       true,
		}

		if !validPaths[r.URL.Path] {
			ErrorPage(w, r, http.StatusNotFound, "Artist page Not Found")
			return
		}

		if r.Method != "GET" {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		artists, err := fetchArtistsFunc()
		if err != nil {
			ErrorPage(w, r, http.StatusBadRequest, "Error fetching artist data")
			return
		}

		execErr := tmpl.ExecuteTemplate(w, "index.html", artists)
		if execErr != nil {
			ErrorPage(w, r, http.StatusInternalServerError, "Error rendering index template")
		}
	}
}

func TestArtistHandler(t *testing.T) {
	// Setup the template parsing for the test
	tmpl, _ = template.New("index.html").Parse("<html><body>{{.Name}}</body></html>")

	// Define test cases
	tests := []struct {
		name       string
		urlPath    string
		method     string
		wantStatus int
	}{
		{
			name:       "Valid path and GET method",
			urlPath:    "/artistFullProfile",
			method:     "GET",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid path",
			urlPath:    "/invalidPath",
			method:     "GET",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid method",
			urlPath:    "/artistFullProfile",
			method:     "POST",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	// Create a new HTTP request
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.urlPath, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Use the mock handler with the mock FetchArtists function
			handler := ArtistHandlerWithMock(mockFetchArtists)

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}
}
