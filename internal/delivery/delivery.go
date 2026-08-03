// Package delivery implements synchoronous delivery logic.
package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/droffilc1/webhook-service/internal/store"
	"github.com/google/uuid"
)

type DeliveryService struct {
	store  store.Store
	client *http.Client
}

func New(store store.Store, client *http.Client) *DeliveryService {
	if client == nil {
		client = http.DefaultClient
	}

	return &DeliveryService{
		store:  store,
		client: client,
	}
}

func (d *DeliveryService) DeliverEvent(event *model.Event) error {
	// Record the event.
	if err := d.store.CreateEvent(event); err != nil {
		return err
	}

	endpoints, err := d.store.ListEndpoints()
	if err != nil {
		return err
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	var errs []error

	for _, endpoint := range endpoints {
		if err := d.deliverToEndpoint(event, endpoint, data); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (d *DeliveryService) deliverToEndpoint(
	event *model.Event,
	endpoint *model.Endpoint,
	data []byte,
) error {

	delivery := &model.Delivery{
		ID:         uuid.NewString(),
		EventID:    event.ID,
		EndpointID: endpoint.ID,
		Attempt:    1,
		CreatedAt:  time.Now(),
	}

	resp, err := d.client.Post(
		endpoint.URL,
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		delivery.Status = model.DeliveryStatusFailed
		_ = d.store.CreateDelivery(delivery)

		return err
	}

	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	delivery.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		delivery.Status = model.DeliveryStatusSuccess
	} else {
		delivery.Status = model.DeliveryStatusFailed
	}

	if err := d.store.CreateDelivery(delivery); err != nil {
		return err
	}

	if delivery.Status == model.DeliveryStatusFailed {
		return fmt.Errorf("endpoint returned %d", resp.StatusCode)
	}

	return nil
}
