package mocks

import (
	model "github.com/dafian47/go-freya-rest-api/module/event"
	"github.com/stretchr/testify/mock"
	base "gitlab.com/xeranta/ustadz-stream/ustadz-stream-go/module"
	"log"
)

var (
	dataList = []model.Event{
		{ID: "01", Name: "Name01", Location: "Location01"},
		{ID: "02", Name: "Name02", Location: "Location02"},
		{ID: "03", Name: "Name03", Location: "Location03"},
		{ID: "04", Name: "Name04", Location: "Location04"},
		{ID: "05", Name: "Name05", Location: "Location05"},
	}
)

type EventRepository struct {
	mock.Mock
}

func (r *EventRepository) GetEventAll() ([]model.Event, error) {
	return dataList, nil
}

func (r *EventRepository) GetEvent(eventID string) (model.Event, error) {

	log.Println("ID " + eventID)

	var data model.Event

	for i := range dataList {

		log.Println("Data List " + dataList[i].ID)

		if dataList[i].ID == eventID {
			data = dataList[i]
			break
		}
	}

	if data.ID == "" {
		return data, base.NOT_FOUND_ERROR
	}

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
