package inmemory

import (
	"github.com/lenniDespero/otus-golang/hw13/internal/event"
	stor "github.com/lenniDespero/otus-golang/hw13/internal/storage"
	"reflect"
	"testing"
	"time"
)

func prepareStorage() *Storage {
	storage := New()

	storage.events = map[int64]event.Event{
		1: {
			ID:           1,
			Title:        "first title",
			DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
			DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
			EditorID:     12,
			CreatorID:    13,
			Deleted:      false,
			DateStarted:  time.Date(2020, time.January, 1, 13, 20, 12, 0, time.Local),
			DateComplete: time.Date(2020, time.January, 1, 14, 20, 12, 0, time.Local)},
		2: {
			ID:           2,
			Title:        "first title",
			DateCreated:  time.Date(2020, time.January, 2, 10, 20, 12, 0, time.Local),
			DateEdited:   time.Date(2020, time.January, 2, 10, 20, 12, 0, time.Local),
			EditorID:     12,
			CreatorID:    13,
			Deleted:      false,
			DateStarted:  time.Date(2020, time.January, 2, 10, 20, 12, 0, time.Local),
			DateComplete: time.Date(2020, time.January, 2, 20, 20, 12, 0, time.Local)},
		3: {
			ID:           3,
			Title:        "first title",
			DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
			DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
			EditorID:     12,
			CreatorID:    13,
			Deleted:      true,
			DateStarted:  time.Date(2020, time.January, 3, 9, 20, 12, 0, time.Local),
			DateComplete: time.Date(2020, time.January, 3, 23, 20, 12, 0, time.Local)},
	}
	return storage
}

func TestNew(t *testing.T) {
	storage1 := New()
	//storage1.events[1] = event.Event{ID: 1, Title: "2"}
	storage2 := New()
	if !reflect.DeepEqual(storage1, storage2) {
		t.Errorf("Not equal data in storage: %v, %v", storage1, storage2)
	}
}

func TestStorage_Add(t *testing.T) {
	storage := prepareStorage()
	newEvent := event.Event{
		ID:           4,
		Title:        "first title",
		DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		EditorID:     12,
		CreatorID:    13,
		Deleted:      true,
		DateStarted:  time.Date(2020, time.January, 2, 21, 20, 12, 0, time.Local),
		DateComplete: time.Date(2020, time.January, 2, 22, 20, 12, 0, time.Local)}
	badEvent := event.Event{
		ID:           5,
		Title:        "first title",
		DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		EditorID:     12,
		CreatorID:    13,
		Deleted:      true,
		DateStarted:  time.Date(2020, time.January, 2, 21, 20, 12, 0, time.Local),
		DateComplete: time.Date(2020, time.January, 3, 10, 20, 12, 0, time.Local)}
	err := storage.Add(newEvent)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	err = storage.Add(badEvent)
	if err == nil {
		t.Errorf("expected error: %s, but get nil", stor.ErrDateBusy)
	}
}

func TestStorage_Delete(t *testing.T) {
	storage := prepareStorage()
	err := storage.Delete(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	err = storage.Delete(1)
	if err == nil {
		t.Errorf("expected error: %s", stor.ErrEventDeleted)
	} else if err != stor.ErrEventDeleted {
		t.Errorf("expected error: %s, get %s", stor.ErrEventDeleted, err.Error())
	}
	err = storage.Delete(4)
	if err == nil {
		t.Errorf("expected error: %s", stor.ErrNotFound)
	} else if err != stor.ErrNotFound {
		t.Errorf("expected error: %s, get : %s", stor.ErrNotFound, err.Error())
	}
}

func TestStorage_Edit(t *testing.T) {
	//testEvent := event.Event{
	//	ID:           1,
	//	Title:        "first title",
	//	DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
	//	DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
	//	EditorID:     12,
	//	CreatorID:    13,
	//	Deleted:      false,
	//	DateStarted:  time.Date(2020, time.January, 1, 13, 20, 12, 0, time.Local),
	//	DateComplete: time.Date(2020, time.January, 1, 14, 20, 12, 0, time.Local)}
}

func TestStorage_GetEventByID(t *testing.T) {
	testEvent := event.Event{
		ID:           1,
		Title:        "first title",
		DateCreated:  time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		DateEdited:   time.Date(2020, time.January, 1, 12, 20, 12, 0, time.Local),
		EditorID:     12,
		CreatorID:    13,
		Deleted:      false,
		DateStarted:  time.Date(2020, time.January, 1, 13, 20, 12, 0, time.Local),
		DateComplete: time.Date(2020, time.January, 1, 14, 20, 12, 0, time.Local)}
	storage := prepareStorage()
	e, err := storage.GetEventByID(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if !reflect.DeepEqual(e[0], testEvent) {
		t.Errorf("not equal events: %v, %v", e[0], testEvent)
	}
	err = storage.Delete(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	_, err = storage.GetEventByID(1)
	if err == nil {
		t.Errorf("expected error: %s, get nil", stor.ErrEventDeleted)
	} else if err != stor.ErrEventDeleted {
		t.Errorf("expected error: %s, get: %s", stor.ErrEventDeleted, err.Error())
	}
	_, err = storage.GetEventByID(13)
	if err == nil {
		t.Errorf("expected error: %s, get nil", stor.ErrEventDeleted)
	} else if err != stor.ErrNotFound {
		t.Errorf("expected error: %s, get: %s", stor.ErrNotFound, err.Error())
	}
}

func TestStorage_GetEvents(t *testing.T) {

}

func Test_inTimeSpan(t *testing.T) {
	test := []struct {
		start  string
		end    string
		check  string
		isTrue bool
	}{
		{"23:00", "05:00", "04:00", true},
		{"23:00", "05:00", "23:30", true},
		{"23:00", "05:00", "20:00", false},
		{"10:00", "21:00", "11:00", true},
		{"10:00", "21:00", "22:00", false},
		{"10:00", "21:00", "03:00", false},
		{"22:00", "02:00", "00:00", true},
		{"10:00", "21:00", "10:00", true},
		{"10:00", "21:00", "21:00", true},
		{"23:00", "05:00", "06:00", false},
		{"23:00", "05:00", "23:00", true},
		{"23:00", "05:00", "05:00", true},
		{"10:00", "21:00", "10:00", true},
		{"10:00", "21:00", "21:00", true},
		{"10:00", "10:00", "09:00", false},
		{"10:00", "10:00", "11:00", false},
		{"10:00", "10:00", "10:00", true},
	}
	newLayout := "15:04"
	for _, row := range test {
		check, _ := time.Parse(newLayout, row.check)
		start, _ := time.Parse(newLayout, row.start)
		end, _ := time.Parse(newLayout, row.end)
		result := inTimeSpan(start, end, check)
		if result != row.isTrue {
			t.Errorf("get %t, expected %t on row: {%s, %s, %s, %t}", result, row.isTrue, row.start, row.end, row.check, row.isTrue)
		}
	}
}
