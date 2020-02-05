package storage

import (
	"time"

	stor "github.com/lenniDespero/otus-golang/hw13/internal/models"
	"github.com/lenniDespero/otus-golang/hw13/internal/types"
)

// Storage struct
type Storage struct {
	events map[int64]types.Event
}

var maxId int64 = 0

//New returns new storage
func New() *Storage {
	return &Storage{events: make(map[int64]types.Event)}
}

// Add types to storage.
func (storage *Storage) Add(event types.Event) (int64, error) {
	for _, e := range storage.events {
		if inTimeSpan(e.DateStarted, e.DateComplete, event.DateStarted) ||
			inTimeSpan(e.DateStarted, e.DateComplete, event.DateComplete) ||
			inTimeSpan(event.DateStarted, event.DateComplete, e.DateStarted) ||
			inTimeSpan(event.DateStarted, event.DateComplete, e.DateComplete) {
			return 0, stor.ErrDateBusy
		}
	}
	if event.ID == 0 {
		maxId++
		event.ID = maxId
	}
	_, ok := storage.events[event.ID]
	if ok {
		return 0, stor.ErrEventIdExists
	}
	storage.events[event.ID] = event
	return event.ID, nil
}

// Edit types data in data storage
func (storage *Storage) Edit(id int64, event types.Event) error {
	e, ok := storage.events[id]
	if !ok {
		return stor.ErrNotFound
	} else if e.Deleted == true {
		return stor.ErrEventDeleted
	}
	if event.ID != id {
		delete(storage.events, id)
	}
	storage.events[event.ID] = event
	return nil
}

// GetEvents return all events
func (storage *Storage) GetEvents() ([]types.Event, error) {
	if len(storage.events) > 0 {
		events := make([]types.Event, 0, len(storage.events))
		for _, e := range storage.events {
			if !e.Deleted {
				events = append(events, e)
			}
		}
		if len(events) > 0 {
			return events, nil
		}
	}
	return []types.Event{}, stor.ErrNotFound
}

//GetEventByID return types with ID
func (storage *Storage) GetEventByID(id int64) ([]types.Event, error) {
	e, ok := storage.events[id]
	if !ok {
		return []types.Event{}, stor.ErrNotFound
	} else if e.Deleted == true {
		return []types.Event{}, stor.ErrEventDeleted
	}
	return []types.Event{e}, nil
}

//Delete will mark types as deleted
func (storage *Storage) Delete(id int64) error {
	e, ok := storage.events[id]
	if !ok {
		return stor.ErrNotFound
	} else if e.Deleted == true {
		return stor.ErrEventDeleted
	}
	e.Deleted = true
	storage.events[id] = e
	return nil
}

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
