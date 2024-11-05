package main

import (
	"fmt"
	"groupie/route"
	"net/http"
)

func main() {
	route.Routes()
	server := http.Server{
		Addr : "127.0.0.1:8000",
	}
	fmt.Println("Server running at http://localHost:8000")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Failed to start the server", err)
	}
}


