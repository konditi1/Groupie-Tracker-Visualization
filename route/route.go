package route

import (
	"net/http"

	"groupie/handlers"
)

func Routes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", handlers.ArtistHandler)
	http.HandleFunc("/locations", handlers.LocationHandler)
	http.HandleFunc("/dates", handlers.DatesHandler)
	http.HandleFunc("/relations", handlers.RelationHandler)
	http.HandleFunc("/artistFullProfile", handlers.ArtistFullProfileHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
}
