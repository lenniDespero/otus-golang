package inmemory

import (
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/event"
	stor "github.com/lenniDespero/otus-golang/hw13/internal/storage"
)

// Storage struct
type Storage struct {
	events map[int64]event.Event
}

func New() *Storage {
	return &Storage{events: make(map[int64]event.Event)}
}

func (storage *Storage) Add(event event.Event) error {
	for _, e := range storage.events {
		if inTimeSpan(e.DateStarted, e.DateComplete, event.DateStarted) ||
			inTimeSpan(e.DateStarted, e.DateComplete, event.DateComplete) ||
			inTimeSpan(event.DateStarted, event.DateComplete, e.DateStarted) ||
			inTimeSpan(event.DateStarted, event.DateComplete, e.DateComplete) {
			return stor.ErrDateBusy
		}
	}
	storage.events[event.ID] = event
	return nil
}

func (storage *Storage) Edit(id int64, event event.Event) error {
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

func (storage *Storage) GetEvents() ([]event.Event, error) {
	if len(storage.events) > 0 {
		events := make([]event.Event, 0, len(storage.events))
		for _, e := range storage.events {
			if !e.Deleted {
				events = append(events, e)
			}
		}
		if len(events) > 0 {
			return events, nil
		}
	}
	return []event.Event{}, stor.ErrNotFound
}

//GetEventByID
func (storage *Storage) GetEventByID(id int64) ([]event.Event, error) {
	e, ok := storage.events[id]
	if !ok {
		return []event.Event{}, stor.ErrNotFound
	} else if e.Deleted == true {
		return []event.Event{}, stor.ErrEventDeleted
	}
	return []event.Event{e}, nil
}

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
