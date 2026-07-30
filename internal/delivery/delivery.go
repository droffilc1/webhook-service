// Package delivery implements synchoronous delivery logic.
package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/droffilc1/webhook-service/internal/store"
)

type DeliveryService struct {
	store  store.Store
	client *http.Client
}

func (d *DeliveryService) DeliveryEvent(event *model.Event) error {
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
	for _, endpoint := range endpoints {
		resp, err := d.client.Post(
			endpoint.URL,
			"application/json",
			bytes.NewBuffer(data))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	return nil
}
