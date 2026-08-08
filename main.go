package main

import (
	"fmt"
	"net/http"
)

// Init server & attaching everything
func main() {
	server := NewServer()

	http.Handle("/", http.FileServer(http.Dir("web")))

	// Attaching route
	http.HandleFunc("/send", server.handleSend)

	// Attaching msgs
	http.HandleFunc("/messages", server.handleMessages)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
