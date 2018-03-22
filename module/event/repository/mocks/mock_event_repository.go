package mocks

import (
	model "github.com/dafian47/go-freya-rest-api/module/event"
	"github.com/stretchr/testify/mock"
)

type EventRepository struct {
	mock.Mock
}

func (r *EventRepository) GetEventAll() ([]model.Event, error) {

	d1 := model.Event{ID: "01", Name: "Name01", Location: "Location01"}
	d2 := model.Event{ID: "02", Name: "Name02", Location: "Location02"}
	d3 := model.Event{ID: "03", Name: "Name03", Location: "Location03"}

	dataList := []model.Event{d1, d2, d3}

	return dataList, nil
}

func (r *EventRepository) GetEvent(eventID string) (model.Event, error) {

	args := r.Called(eventID)
	data := args.Get(0).(model.Event)

	return data, nil
}

func (r *EventRepository) CreateEvent(event model.Event) (model.Event, error) {

	args := r.Called(event)
	data := args.Get(0).(model.Event)

	return data, nil
}

func (r *EventRepository) UpdateEvent(event model.Event) (model.Event, error) {

	args := r.Called(event)
	data := args.Get(0).(model.Event)

	return data, nil
}

func (r *EventRepository) DeleteEvent(eventID string) (bool, error) {

	args := r.Called(eventID)
	data := args.Get(0).(model.Event)

	if data.ID == "" {
		return false, nil
	}

	return true, nil
}
