package storage

import (
	"github.com/lenniDespero/otus-golang/hw13/internal/event"
)

// Storage interface
type Storage interface {
	// Add event to storage.
	Add(event event.Event) error

	// Remove event from data storage.
	Remove(id int64) error

	// Edit event data in data storage
	Edit(id int64, event event.Event) error

	// GetEvents return all events
	GetEvents() ([]event.Event, error)

	//GetEventsBy return events with params
	GetEventsBy(property string, value string) ([]event.Event, error)
}
