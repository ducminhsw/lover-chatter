package model

import "time"

type Conversation struct {
	Id      string    `json:"id"`   // index 1
	Page    int       `json:"page"` // index 2
	From    string    `json:"from"`
	To      string    `json:"to"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sentat"`
}
