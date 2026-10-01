package model

import "time"

// APIKey describes properties of APIKey
type APIKey struct {
	ID        string     `json:"id"`
	KeyHash   string     `json:"key_hash"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}
