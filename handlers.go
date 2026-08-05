package main

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleSend(w http.ResponseWriter, r *http.Request) {

	var msg Message

	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// // Decode incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Adding the message to memory
	s.AddMessage(msg)

	// All good code 200
	w.WriteHeader(http.StatusOK)
}
