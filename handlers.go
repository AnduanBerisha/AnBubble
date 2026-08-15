package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"unicode/utf8"
)

func (s *Server) handleSend(w http.ResponseWriter, r *http.Request) {

	var msg Message

	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// limits size of body to avoid spam
	r.Body = http.MaxBytesReader(w, r.Body, 64*1024)

	// Decode incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	msg.Sender = strings.TrimSpace(msg.Sender)
	msg.Text = strings.TrimSpace(msg.Text)

	if msg.Sender == "" || utf8.RuneCountInString(msg.Sender) > 25 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if msg.Text == "" || utf8.RuneCountInString(msg.Text) > 2000 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	msg.ID = generateID()

	// Adding the message to memory
	s.AddMessage(msg)

	// All good code 200
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleMessages(w http.ResponseWriter, r *http.Request) {

	// Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	messages := s.GetMessages()

	json.NewEncoder(w).Encode(messages)
}

func generateID() string {
	bytes := make([]byte, 8) // Allocate an 8 byte buffer

	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
