package calendar

import (
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/types"
)

// CalendarInterface interface
type CalendarInterface interface {
	//New will create new calendar
	New(storage types.StorageInterface) CalendarInterface

	// Add event to calendar.
	Add(title string, dateStarted time.Time, dateComplete time.Time, notice string, userId int64) (int64, error)

	//Edit event data in calendar
	Edit(id string, event models.Event, userId int64) error

	//GetEvents return all events
	GetEvents() ([]models.Event, error)

	//GetEventByID return event with ID
	GetEventByID(id string) ([]models.Event, error)

	//Delete will mark event as deleted
	Delete(id string) error
}
