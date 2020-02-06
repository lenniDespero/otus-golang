package calendar

import (
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/types"
)

type Calendar struct {
	storage types.StorageInterface
}

//New will create new calendar
func New(storage types.StorageInterface) *Calendar {
	return &Calendar{storage: storage}
}

// Add event to calendar.
func (c Calendar) Add(title string, dateStarted time.Time, dateComplete time.Time, notice string, userId int64) (string, error) {
	event := models.Event{
		ID:           "",
		Title:        title,
		DateEdited:   time.Now(),
		EditorID:     userId,
		DateCreated:  time.Now(),
		CreatorID:    userId,
		DateStarted:  dateStarted,
		DateComplete: dateComplete,
		Notice:       notice,
		Deleted:      false,
	}
	return c.storage.Add(event)
}

//Edit event data in calendar
func (c Calendar) Edit(id string, event models.Event, userId int64) error {
	event.EditorID = userId
	event.DateEdited = time.Now()
	return c.storage.Edit(id, event)
}

//GetEvents return all events
func (c Calendar) GetEvents() ([]models.Event, error) {
	return c.storage.GetEvents()
}

//GetEventByID return event with ID
func (c Calendar) GetEventByID(id string) ([]models.Event, error) {
	return c.storage.GetEventByID(id)
}

//Delete will mark event as deleted
func (c Calendar) Delete(id string) error {
	return c.storage.Delete(id)
}
