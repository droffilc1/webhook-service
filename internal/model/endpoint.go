package model

import "time"

// Endpoint describes properties of an endpoint
type Endpoint struct {
	ID        string    `json:"id"`
	URL       string    `json:"string"`
	CreatedAt time.Time `json:"created_at"`
}
