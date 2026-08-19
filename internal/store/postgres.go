package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/droffilc1/webhook-service/internal/model"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) Store {
	return &PostgresStore{db: db}
}

// CreateDelivery implements [Store].
func (p *PostgresStore) CreateDelivery(delivery *model.Delivery) error {
	query := `
	INSERT INTO deliveries (id, event_id,endpoint_id, status, attempt, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query,
		delivery.ID,
		delivery.EventID,
		delivery.EndpointID,
		delivery.Status,
		delivery.Attempt,
		delivery.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// CreateEndpoint implements [Store].
func (p *PostgresStore) CreateEndpoint(endpoint *model.Endpoint) error {
	query := `
	INSERT INTO endpoints (id, url, created_at)
	VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(query,
		endpoint.ID,
		endpoint.URL,
		endpoint.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// CreateEvent implements [Store].
func (p *PostgresStore) CreateEvent(event *model.Event) error {
	query := `
	INSERT INTO events (id, type, payload, created_at)
	VALUES ($1, $2, $3, 4$)
	`
	_, err := p.db.Exec(query,
		event.ID,
		event.Type,
		event.Payload,
		event.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEndpoint implements [Store].
func (p *PostgresStore) DeleteEndpoint(id string) error {
	query := `DELETE FROM endpoints WHERE id = $1`

	result, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrEndpointNotFound
	}

	return nil
}

// GetDelivery implements [Store].
func (p *PostgresStore) GetDelivery(id string) (*model.Delivery, error) {
	query := `
	SELECT id, event_id, endpoint_id, status, attempt, created_at
	FROM deliveries
	WHERE id = $1
	`

	var d model.Delivery
	err := p.db.QueryRow(query, id).Scan(
		&d.ID,
		&d.EventID,
		&d.EndpointID,
		&d.Status,
		&d.Attempt,
		&d.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDeliveryNotFound
		}

		return nil, err
	}
	return &d, nil
}

// GetEndpoint implements [Store].
func (p *PostgresStore) GetEndpoint(id string) (*model.Endpoint, error) {
	query := `
	SELECT id, url, created_at
	FROM endpoints
	WHERE id = $1
	`
	var e model.Endpoint
	err := p.db.QueryRow(query, id).Scan(
		&e.ID,
		&e.URL,
		&e.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEndpointNotFound
		}
		return nil, err
	}
	return &e, nil
}

// GetEvent implements [Store].
func (p *PostgresStore) GetEvent(id string) (*model.Event, error) {
	query := `
	SELECT id, type, payload, created_at
	FROM events
	WHERE id = $1
	`

	var ev model.Event
	err := p.db.QueryRow(query, id).Scan(
		&ev.ID,
		&ev.Type,
		&ev.Payload,
		&ev.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEventNotFound
		}
		return nil, err
	}
	return &ev, nil
}

// ListDeliveries implements [Store].
func (p *PostgresStore) ListDeliveries() ([]*model.Delivery, error) {
	query := `
	SELECT id, event_id, endpoint_id, status, attempt, created_at
	FROM deliveries
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deliveries []*model.Delivery
	for rows.Next() {
		var d model.Delivery
		err = rows.Scan(
			&d.ID,
			&d.EventID,
			&d.EndpointID,
			&d.Status,
			&d.Attempt,
			&d.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, &d)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return deliveries, nil
}

// ListEndpoints implements [Store].
func (p *PostgresStore) ListEndpoints() ([]*model.Endpoint, error) {
	query := `
	SELECT id, url, created_at
	FROM endpoints
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var endpoints []*model.Endpoint
	for rows.Next() {
		var e model.Endpoint
		err = rows.Scan(
			&e.ID,
			&e.URL,
			&e.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, &e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return endpoints, nil
}

// ListEvents implements [Store].
func (p *PostgresStore) ListEvents() ([]*model.Event, error) {
	query := `
	SELECT id, type, payload, created_at
	FROM events
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*model.Event
	for rows.Next() {
		var ev model.Event
		err = rows.Scan(
			&ev.ID,
			&ev.Type,
			&ev.Payload,
			&ev.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &ev)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// UpdateEndpoint implements [Store].
func (p *PostgresStore) UpdateEndpoint(endpoint *model.Endpoint) error {
	query := `
	UPDATE endpoints
	SET url = $1, updated_at = $2, 
	WHERE id = $3
	`

	now := time.Now()
	result, err := p.db.Exec(query, endpoint.URL, now, endpoint.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrEndpointNotFound
	}

	endpoint.UpdatedAt = now
	return nil
}
