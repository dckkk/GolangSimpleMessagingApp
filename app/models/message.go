package models

import "time"

type MessagePayload struct {
	From    string    `json:"from"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}
