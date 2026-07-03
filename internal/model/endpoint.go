package model

import "time"

// Endpoint describes properties of an endpoint
type Endpoint struct {
	ID      string
	URL     string
	Created time.Time
}
