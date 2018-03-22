package repository

import (
	"errors"
	model "github.com/dafian47/go-freya-rest-api/module/event"
	"github.com/jinzhu/gorm"
)

type EventRepository interface {
	GetEventAll() ([]model.Event, error)
	GetEvent(eventID string) (model.Event, error)
	CreateEvent(event model.Event) (model.Event, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEvent(eventID string) (bool, error)
}

type eventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepository {
	return &eventRepo{DB: db}
}

func (r *eventRepo) GetEventAll() ([]model.Event, error) {

	var itemList []model.Event

	r.DB.Find(&itemList)

	if len(itemList) == 0 {
		return nil, errors.New("empty data")
	}

	return itemList, nil
}

func (r *eventRepo) GetEvent(eventID string) (model.Event, error) {

	var item model.Event

	r.DB.Where(&model.Event{ID: eventID}).First(&item)

	if item.ID == "" {
		return item, errors.New("not found")
	}

	return item, nil
}

func (r *eventRepo) CreateEvent(event model.Event) (model.Event, error) {

	r.DB.Save(&event)

	if event.ID == "" {
		return event, errors.New("not found")
	}

	return event, nil
}

func (r *eventRepo) UpdateEvent(event model.Event) (model.Event, error) {

	r.DB.Save(&event)

	if event.ID == "" {
		return event, errors.New("not found")
	}

	return event, nil
}

func (r *eventRepo) DeleteEvent(eventID string) (bool, error) {

	var item model.Event

	r.DB.Where(&model.Event{ID: eventID}).First(&item)

	if item.ID == "" {
		return false, errors.New("not found")
	}

	r.DB.Delete(&item)

	return true, nil
}
