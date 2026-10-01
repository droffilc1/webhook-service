package store

import (
	"testing"

	"github.com/droffilc1/webhook-service/internal/model"
)

func TestMemory_CreateAndGetAPIKey(t *testing.T) {
	store := NewStore()

	apiKey := &model.APIKey{
		ID:      "some-id",
		KeyHash: "some-hash",
		Name:    "test",
	}

	err := store.CreateAPIKey(apiKey)
	if err != nil {
		t.Fatal(err)
	}

	got, err := store.GetAPIKeyByHash("some-hash")
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != apiKey.ID {
		t.Errorf("got ID %q, want %q", got.ID, apiKey.ID)
	}
}
