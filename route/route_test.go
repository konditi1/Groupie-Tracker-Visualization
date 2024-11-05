package route

import (
	"net/http"
	"net/http/httptest"
	"testing"

)

// Setup a new mux and register routes
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	Routes() // Ensure this sets up routes with the correct mux
	return mux
}

// Helper function to set up routes
func setupTestServer() *httptest.Server {
	mux := setupRoutes()
	return httptest.NewServer(mux)
}

func TestRoutes(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Define test cases
	tests := []struct {
		path           string
		expectedStatus int
	}{
		{path: "/index", expectedStatus: http.StatusNotFound},
		{path: "/location", expectedStatus: http.StatusNotFound},
		{path: "/date", expectedStatus: http.StatusNotFound},
		{path: "/relation", expectedStatus: http.StatusNotFound},
		{path: "/artistFullProfiles", expectedStatus: http.StatusNotFound},
		{path: "/searchs", expectedStatus: http.StatusNotFound},
		{path: "/static/testfile.txt", expectedStatus: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			resp, err := http.Get(server.URL + tt.path)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("For path %s, expected status code %d but got %d", tt.path, tt.expectedStatus, resp.StatusCode)
			}
		})
	}
}
