package types

import (
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
)

// StorageInterface interface
type StorageInterface interface {

	// Add models to storage.
	Add(event models.Event) (string, error)

	// Edit models data in data storage
	Edit(id string, event models.Event) error

	// GetEvents return all events
	GetEvents() ([]models.Event, error)

	//GetEventByID return models with ID
	GetEventByID(id string) ([]models.Event, error)

	//Delete will mark models as deleted
	Delete(id string) error
}
