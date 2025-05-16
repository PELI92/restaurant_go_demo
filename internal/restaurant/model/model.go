package model

import "time"

type Restaurant struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"url" bson:"url"`
	Description string    `json:"description" bson:"description"`
	Tags        []string  `json:"tags" bson:"tags"`
	Address     string    `json:"interval_in_seconds" bson:"interval_in_seconds"`
	Status      bool      `json:"is_active" bson:"is_active"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
