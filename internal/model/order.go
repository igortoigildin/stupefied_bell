package model

import "time"

type Order struct {
	Number     string    `json:"number"`
	Quantity   int       `json:"quantity"`
	Title      string    `json:"title"`
	UploadedAt time.Time `json:"date"`
	Comment    string    `json:"comment,omitempty"`
	Status     string    `json:"status"`
}
