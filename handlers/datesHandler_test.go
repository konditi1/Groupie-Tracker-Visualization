package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"groupie/models"
	"html/template"
	"fmt"
	"strconv"
)

// Mock function for fetching dates
func mockFetchDates() ([]models.Dates, error) {
	return []models.Dates{
		{Id: 1, Dates: []string{"2024-09-01"}},
		{Id: 2, Dates: []string{"2024-09-02"}},
	}, nil
}

// Mock data for artists
var mockArtists = []models.Artist{
	{Name: "Queen"},
	{Name: "Beatles"},
}

// Mock function for testing
func DatesHandlerWithMock(fetchDatesFunc func() ([]models.Dates, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			ErrorPage(w, r, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		artistIdStr := r.URL.Query().Get("artistId")
		artistId, err := strconv.Atoi(artistIdStr)
		if err != nil {
			ErrorPage(w, r, http.StatusBadRequest, "Invalid artist ID")
			return
		}

		if artistId < 1 || artistId > 52 {
			ErrorPage(w, r, http.StatusBadRequest, "Invalid artist Id provided")
			fmt.Println("Invalid artist Id provided")
			return
		}

		artistName := r.URL.Query().Get("name")
		check := false
		id := 1
		for i, name := range mockArtists {
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

		dates, err := fetchDatesFunc()
		if err != nil {
			ErrorPage(w, r, http.StatusInternalServerError, "Error fetching dates data")
			return
		}

		date := dates[artistId-1]

		DatesData := struct {
			Dates models.Dates
			Name  string
		}{
			Dates: date,
			Name:  artistName,
		}

		execErr := tmpl.ExecuteTemplate(w, "dates.html", DatesData)
		if execErr != nil {
			ErrorPage(w, r, http.StatusInternalServerError, "Error rendering dates template")
		}
	}
}

func TestDatesHandler(t *testing.T) {
	// Setup template parsing for testing
	tmpl, _ = template.New("dates.html").Parse("<html><body>{{.Name}} - {{range .Dates.Dates}}{{.}} {{end}}</body></html>")

	tests := []struct {
		name       string
		urlPath    string
		method     string
		wantStatus int
	}{
		{
			name:       "Valid artistId and name",
			urlPath:    "/dates?artistId=1&name=Queen",
			method:     "GET",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid artistId",
			urlPath:    "/dates?artistId=100&name=Queen",
			method:     "GET",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Invalid method",
			urlPath:    "/dates?artistId=1&name=Queen",
			method:     "POST",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.urlPath, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Use the mock handler with the mock FetchDates function
			handler := DatesHandlerWithMock(mockFetchDates)

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}
}
