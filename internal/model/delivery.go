package model

import "time"

// Delivery describes properties of a delivery
type Delivery struct {
	ID         string
	EventID    string
	EndpointID string
	Status     string
	StatusCode int
	Attempt    int
	Created    time.Time
}
