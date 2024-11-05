package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie/api"
	"groupie/models"
)

// Mock data matching the structure of models.IndexDates
var mockIndexDates = models.IndexDates{
	Index: []models.Dates{
		{
			Id:    1,
			Dates: []string{"2024-01-01", "2024-01-02"}, // Example dates
		},
	},
}

// TestFetchDates checks if FetchDates function fetches data correctly
func TestFetchDates(t *testing.T) {
	// Create a test server with mock data
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockIndexDates)
	}))
	defer testServer.Close()

	// Call the FetchDates function with the test server URL
	dates, err := api.FetchDates()
	if err != nil {
		t.Fatalf("FetchDates() error = %v, want nil", err)
	}

	if len(dates.Index) != len(mockIndexDates.Index) {
		t.Errorf("FetchDates() = %v, want %v", len(dates.Index), len(mockIndexDates.Index))
	}

	// Example check for specific fields; update based on your actual struct
	if dates.Index[0].Id != mockIndexDates.Index[0].Id {
		t.Errorf("FetchDates() = %v, want %v", dates.Index[0].Id, mockIndexDates.Index[0].Id)
	}
	if len(dates.Index[0].Dates) != len(mockIndexDates.Index[0].Dates) {
		t.Errorf("FetchDates() = %v, want %v", len(dates.Index[0].Dates), len(mockIndexDates.Index[0].Dates))
	}
}

// TestFetchDatesRequestError simulates a scenario where the server request fails
func TestFetchDatesRequestError(t *testing.T) {
	// Create a test server that returns a server error
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer testServer.Close()

	// Attempt to fetch data, expecting an error
	_, err := api.FetchDates()

	if err == nil {
		t.Fatalf("FetchDates() error = nil, want error")
	}
}

// TestFetchDatesDecodeError tests the scenario where the server returns invalid JSON
func TestFetchDatesDecodeError(t *testing.T) {
	// Create a test server that returns invalid JSON
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid json")
	}))
	defer testServer.Close()

	// Attempt to fetch data, expecting a decoding error
	_, err := api.FetchDates()

	if err == nil {
		t.Fatalf("FetchDates() error = nil, want error")
	}
}
