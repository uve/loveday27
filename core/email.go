package core

import (
	"appengine/datastore"
	"time"
)

// Score is an entity to store campaign
type Email struct {
	Name  string
	Content string `datastore:",noindex"`

	Status *datastore.Key
	Campaign *datastore.Key
	Created  time.Time
	Modified  time.Time
}
