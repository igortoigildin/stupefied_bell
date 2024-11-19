package model

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Order struct {
	Id         string    `json:"number"`
	Quantity   int       `json:"quantity"`
	Title      string    `json:"title"`
	UploadedAt time.Time `json:"date"`
	Comment    string    `json:"comment,omitempty"`
	Status     string    `json:"status"`
}
