// Package model implements structures of items
package model

import "time"

type DeliveryStatus string

const (
	DeliveryStatusSuccess DeliveryStatus = "success"
	DeliveryStatusFailed  DeliveryStatus = "failed"
)

// Delivery describes properties of a delivery
type Delivery struct {
	ID         string         `json:"id"`
	EventID    string         `json:"event_id"`
	EndpointID string         `json:"endpoint_id"`
	Status     DeliveryStatus `json:"status"`
	StatusCode int            `json:"status_code"`
	Attempt    int            `json:"attempt"`
	CreatedAt  time.Time      `json:"created_at"`
}
