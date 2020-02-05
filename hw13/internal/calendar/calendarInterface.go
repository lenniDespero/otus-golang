package calendar

import (
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/types"
)

// CalendarInterface interface
type CalendarInterface interface {
	//New will create new calendar
	New(storage models.StorageInterface) CalendarInterface

	// Add event to calendar.
	Add(title string, dateStarted time.Time, dateComplete time.Time, notice string, userId int64) (int64, error)

	//Edit event data in calendar
	Edit(id int64, event types.Event, userId int64) error

	//GetEvents return all events
	GetEvents() ([]types.Event, error)

	//GetEventByID return event with ID
	GetEventByID(id int64) ([]types.Event, error)

	//Delete will mark event as deleted
	Delete(id int64) error
}
