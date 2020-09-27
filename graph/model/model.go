package model

import (
	"time"
)

type Message struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
