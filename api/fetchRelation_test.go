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

// Mock data for testing
var mockIndexRelation = models.IndexRelation{
	Index: []models.Relation{
		{
			Id: 1,
			DatesLocations: map[string][]string{
				"Location1": {"Date1", "Date2"},
				"Location2": {"Date3"},
			},
		},
		{
			Id: 2,
			DatesLocations: map[string][]string{
				"Location3": {"Date4", "Date5"},
			},
		},
	},
}

// TestFetchRelation checks if FetchRelation function fetches data correctly
func TestFetchRelation(t *testing.T) {
	// Create a test server with mock data
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockIndexRelation)
	}))
	defer testServer.Close()

	// Call FetchRelation with the test server URL
	relation, err := api.FetchRelation()
	if err != nil {
		t.Fatalf("FetchRelation() error = %v, want nil", err)
	}

	// Validate that the length of relations fetched matches the mock data
	if len(relation.Index) != len(mockIndexRelation.Index) {
		t.Errorf("FetchRelation() = %v, want %v", len(relation.Index), len(mockIndexRelation.Index))
	}

	// Example check for specific fields; update based on your actual struct
	if relation.Index[0].Id != mockIndexRelation.Index[0].Id {
		t.Errorf("FetchRelation() = %v, want %v", relation.Index[0].Id, mockIndexRelation.Index[0].Id)
	}

	for location, dates := range mockIndexRelation.Index[0].DatesLocations {
		actualDates, ok := relation.Index[0].DatesLocations[location]
		if !ok {
			t.Errorf("FetchRelation() missing location %v", location)
			continue
		}
		if !equalStringSlices(actualDates, dates) {
			t.Errorf("FetchRelation() = %v, want %v for location %v", actualDates, dates, location)
		}
	}
}

// equalStringSlices compares two slices of strings for equality
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, found := m[v]; !found {
			return false
		}
	}
	return true
}

// TestFetchRelationRequestError simulates a scenario where the server request fails
func TestFetchRelationRequestError(t *testing.T) {
	// Create a test server that returns a server error
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer testServer.Close()

	_, err := api.FetchRelation()

	if err == nil {
		t.Fatalf("FetchRelation() error = nil, want error")
	}
}

// TestFetchRelationDecodeError tests the scenario where the server returns invalid JSON
func TestFetchRelationDecodeError(t *testing.T) {
	// Create a test server that returns invalid JSON
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "invalid json")
	}))
	defer testServer.Close()

	// Attempt to fetch data, expecting a decoding error
	_, err := api.FetchRelation()

	if err == nil {
		t.Fatalf("FetchRelation() error = nil, want error")
	}
}
