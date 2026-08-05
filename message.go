package main

type Message struct {
	ID     string `json:"id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
