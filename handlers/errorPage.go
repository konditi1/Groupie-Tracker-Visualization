package handlers

import (
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	errorData := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statusCode,
		Message:    message,
	}
	err := tmpl.ExecuteTemplate(w, "error.html", errorData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
