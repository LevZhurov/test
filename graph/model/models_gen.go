// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type NewMessage struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
