package store

import "github.com/droffilc1/webhook-service/internal/model"

// Store defines the properties of data to be stored
type Store interface {
	CreateEvent(event *model.Event) error
	GetEvent(id string) (*model.Event, error)
	ListEvents() ([]*model.Event, error)

	CreateEndpoint(endpoint *model.Endpoint) error
	GetEndpoint(id string) (*model.Endpoint, error)
	ListEndpoints() ([]*model.Endpoint, error)
	UpdateEndpoint(endpoint *model.Endpoint) error
	DeleteEndpoint(id string) error

	CreateDelivery(delivery *model.Delivery) error
	GetDelivery(id string) (*model.Delivery, error)
	ListDeliveries() ([]*model.Delivery, error)
}
