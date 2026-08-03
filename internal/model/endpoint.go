package model

import "time"

// Endpoint describes properties of an endpoint
type Endpoint struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
