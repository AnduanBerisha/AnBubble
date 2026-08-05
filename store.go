package main

import (
	"sync"
)

type Server struct {
	mu       sync.RWMutex // for safely reading & write msgs without crashing when 2+ usrs connects at the same time
	messages []Message    // dynamic array of Messages structures
}

// creates & initializes a new Server instance with empty list of messages
func NewServer() *Server {
	return &Server{messages: make([]Message, 0)} // return memory address (&) of the initialized Server
}

// AddMessage appends new message to memory using a write lock
func (s *Server) AddMessage(msg Message) {
	s.mu.Lock()                          // block server with write only
	defer s.mu.Unlock()                  // auto unlock at the end
	s.messages = append(s.messages, msg) // add the message in the array
}

// GetMessages returns all stored messages using a read lock
func (s *Server) GetMessages() []Message {
	s.mu.RLock()         // block the server with readonly
	defer s.mu.RUnlock() // unlock the reading at the end
	return s.messages    // return msgs
}
