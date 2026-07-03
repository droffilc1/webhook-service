package model

import (
	"encoding/json"
	"time"
)

// Event describes properties of an event
type Event struct {
	ID      string
	Type    string
	Payload json.RawMessage
	Created time.Time
}
