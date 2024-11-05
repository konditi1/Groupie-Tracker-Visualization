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

// Mock data matching the structure of models.IndexLocations
var mockIndexLocations = models.IndexLocations{
	Index: []models.Locations{
		{
			Id:        1,
			Locations: []string{"New York", "Los Angeles"}, // Example locations
		},
	},
}

// TestFetchLocations checks if FetchLocations function fetches data correctly
func TestFetchLocations(t *testing.T) {
	// Create a test server with mock data
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockIndexLocations)
	}))
	defer testServer.Close()

	// Call the FetchLocations function with the test server URL
	locations, err := api.FetchLocations()
	if err != nil {
		t.Fatalf("FetchLocations() error = %v, want nil", err)
	}

	// Validate that the length of locations fetched matches the mock data
	if len(locations.Index) != len(mockIndexLocations.Index) {
		t.Errorf("FetchLocations() = %v, want %v", len(locations.Index), len(mockIndexLocations.Index))
	}

	// Example check for specific fields; update based on your actual struct
	if locations.Index[0].Id != mockIndexLocations.Index[0].Id {
		t.Errorf("FetchLocations() = %v, want %v", locations.Index[0].Id, mockIndexLocations.Index[0].Id)
	}
	if len(locations.Index[0].Locations) != len(mockIndexLocations.Index[0].Locations) {
		t.Errorf("FetchLocations() = %v, want %v", len(locations.Index[0].Locations), len(mockIndexLocations.Index[0].Locations))
	}
}

// TestFetchLocationsRequestError simulates a scenario where the server request fails
func TestFetchLocationsRequestError(t *testing.T) {
	// Create a test server that returns a server error
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer testServer.Close()

	_, err := api.FetchLocations()

	if err == nil {
		t.Fatalf("FetchLocations() error = nil, want error")
	}
}

// TestFetchLocationsDecodeError tests the scenario where the server returns invalid JSON
func TestFetchLocationsDecodeError(t *testing.T) {
	// Create a test server that returns invalid JSON
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid json")
	}))
	defer testServer.Close()

	_, err := api.FetchLocations()

	if err == nil {
		t.Fatalf("FetchLocations() error = nil, want error")
	}
}
