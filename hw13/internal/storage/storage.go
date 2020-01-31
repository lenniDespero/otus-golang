package base

import (
	"github.com/lenniDespero/otus-golang/hw13/internal/event"
)

// StorageInterface interface
type StorageInterface interface {
	// Add event to storage.
	Add(event event.Event) error

	// Remove event from data storage.
	Remove(id int64) error

	// Edit event data in data storage
	Edit(id int64, event event.Event) error

	// GetEvents return all events
	GetEvents() ([]event.Event, error)

	//GetEventByID return event with ID
	GetEventByID(id int64) ([]event.Event, error)

	//Delete will mark event as deleted
	Delete(id int64) error
}
