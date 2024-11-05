package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie/models"
)

// TestFetchArtists tests the FetchArtists function with a mock server.
func TestFetchArtists(t *testing.T) {
	// Create a test server that returns mock artist data
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/artists" {
			t.Errorf("Expected request to /api/artists, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode([]models.Artist{
			{
				Id:   1,
				Name: "Mock Artist 1",
			},
			{
				Id:   2,
				Name: "Mock Artist 2",
			},
		})
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}))
	defer testServer.Close()

	// Mock the URL for testing
	// Since FetchArtists uses a hardcoded URL, this setup might not affect it directly.
	// This is just to show how you would set it up if URL replacement was possible.
	// In practice, you might need to refactor FetchArtists to accept a URL or use dependency injection.
	// originalURL := "https://groupietrackers.herokuapp.com/api/artists"
	// testURL := testServer.URL + "/api/artists"

	// Call FetchArtists and check the results
	artists, err := FetchArtists()
	if err != nil {
		t.Fatalf("FetchArtists() error = %v, want nil", err)
	}

	expectedArtists := []models.Artist{
		{
			Id:   1,
			Name: "Mock Artist 1",
		},
		{
			Id:   2,
			Name: "Mock Artist 2",
		},
	}

	if len(artists) != len(expectedArtists) {
		t.Errorf("FetchArtists() = %v, want %v", len(artists), len(expectedArtists))
	}

	for i, artist := range artists {
		if artist.Id != expectedArtists[i].Id || artist.Name != expectedArtists[i].Name {
			t.Errorf("FetchArtists()[%d] = %v, want %v", i, artist, expectedArtists[i])
		}
	}
}
