// Package store implements storage
package store

import (
	"errors"

	"github.com/droffilc1/webhook-service/internal/model"
)

// Memory keeps the service data in memory
type Memory struct {
	events     map[string]*model.Event
	endpoints  map[string]*model.Endpoint
	deliveries map[string]*model.Delivery
	apiKeys    map[string]*model.APIKey
}

func NewStore() Store {
	return &Memory{
		events:     make(map[string]*model.Event),
		endpoints:  make(map[string]*model.Endpoint),
		deliveries: make(map[string]*model.Delivery),
		apiKeys:    make(map[string]*model.APIKey),
	}
}

var (
	ErrEventNotFound    = errors.New("event not found")
	ErrEndpointNotFound = errors.New("endpoint not found")
	ErrDeliveryNotFound = errors.New("delivery not found")

	ErrEventExists    = errors.New("event already exists")
	ErrEndpointExists = errors.New("endpoint already exists")
	ErrDeliveryExists = errors.New("delivery already exists")

	ErrAPIKeyNotFound = errors.New("API key not found")
)

// CreateAPIKey implements [Store].
func (m *Memory) CreateAPIKey(apiKey *model.APIKey) error {
	m.apiKeys[apiKey.KeyHash] = apiKey
	return nil
}

// GetAPIKeyByHash implements [Store].
func (m *Memory) GetAPIKeyByHash(hash string) (*model.APIKey, error) {
	apiKey, ok := m.apiKeys[hash]
	if !ok {
		return nil, ErrAPIKeyNotFound
	}

	return apiKey, nil
}

// CreateDelivery implements [Store].
func (m *Memory) CreateDelivery(delivery *model.Delivery) error {
	if _, exists := m.deliveries[delivery.ID]; exists {
		return ErrDeliveryExists
	}

	m.deliveries[delivery.ID] = delivery
	return nil
}

// CreateEndpoint implements [Store].
func (m *Memory) CreateEndpoint(endpoint *model.Endpoint) error {
	if _, exists := m.endpoints[endpoint.ID]; exists {
		return ErrEndpointExists
	}
	m.endpoints[endpoint.ID] = endpoint
	return nil
}

// CreateEvent implements [Store].
func (m *Memory) CreateEvent(event *model.Event) error {
	if _, exists := m.events[event.ID]; exists {
		return ErrEventExists
	}
	m.events[event.ID] = event
	return nil
}

// DeleteEndpoint implements [Store].
func (m *Memory) DeleteEndpoint(id string) error {
	if _, ok := m.endpoints[id]; !ok {
		return ErrEndpointNotFound
	}

	delete(m.endpoints, id)
	return nil
}

// GetDelivery implements [Store].
func (m *Memory) GetDelivery(id string) (*model.Delivery, error) {
	delivery, ok := m.deliveries[id]
	if !ok {
		return nil, ErrDeliveryNotFound
	}

	return delivery, nil
}

// GetEndpoint implements [Store].
func (m *Memory) GetEndpoint(id string) (*model.Endpoint, error) {
	endpoint, ok := m.endpoints[id]
	if !ok {
		return nil, ErrEndpointNotFound
	}

	return endpoint, nil
}

// GetEvent implements [Store].
func (m *Memory) GetEvent(id string) (*model.Event, error) {
	event, ok := m.events[id]
	if !ok {
		return nil, ErrEventNotFound
	}
	return event, nil
}

// ListDeliveries implements [Store].
func (m *Memory) ListDeliveries() ([]*model.Delivery, error) {
	deliveries := make([]*model.Delivery, 0, len(m.deliveries))

	for _, delivery := range m.deliveries {
		deliveries = append(deliveries, delivery)
	}

	return deliveries, nil
}

// ListEndpoints implements [Store].
func (m *Memory) ListEndpoints() ([]*model.Endpoint, error) {
	endpoints := make([]*model.Endpoint, 0, len(m.endpoints))

	for _, endpoint := range m.endpoints {
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

// ListEvents implements [Store].
func (m *Memory) ListEvents() ([]*model.Event, error) {
	events := make([]*model.Event, 0, len(m.events))

	for _, event := range m.events {
		events = append(events, event)
	}
	return events, nil
}

// UpdateEndpoint implements [Store].
func (m *Memory) UpdateEndpoint(endpoint *model.Endpoint) error {
	if _, ok := m.endpoints[endpoint.ID]; !ok {
		return ErrEndpointNotFound
	}

	m.endpoints[endpoint.ID] = endpoint
	return nil
}
