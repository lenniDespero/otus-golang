package types

import (
	"context"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
)

// StorageInterface interface
type StorageInterface interface {

	// Add models to storage.
	Add(event models.Event, ctx context.Context) (string, error)

	// Edit models data in data storage
	Edit(id string, event models.Event, ctx context.Context) error

	// GetEvents return all events
	GetEvents(ctx context.Context) ([]models.Event, error)

	//GetEventByID return models with ID
	GetEventByID(id string, ctx context.Context) ([]models.Event, error)

	//Delete will mark models as deleted
	Delete(id string, ctx context.Context) error
}
