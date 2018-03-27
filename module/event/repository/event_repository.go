package repository

import (
	base "github.com/dafian47/go-freya-rest-api/module"
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
		return nil, base.EMPTY_ERROR
	}

	return itemList, nil
}

func (r *eventRepo) GetEvent(eventID string) (model.Event, error) {

	var item model.Event

	r.DB.Where(&model.Event{ID: eventID}).First(&item)

	if item.ID == "" {
		return item, base.NOT_FOUND_ERROR
	}

	return item, nil
}

func (r *eventRepo) CreateEvent(event model.Event) (model.Event, error) {

	r.DB.Save(&event)

	if event.ID == "" {
		return event, base.FAILED_SAVE_ERROR
	}

	return event, nil
}

func (r *eventRepo) UpdateEvent(event model.Event) (model.Event, error) {

	r.DB.Save(&event)

	if event.ID == "" {
		return event, base.FAILED_UPDATE_ERROR
	}

	return event, nil
}

func (r *eventRepo) DeleteEvent(eventID string) (bool, error) {

	var item model.Event

	r.DB.Where(&model.Event{ID: eventID}).First(&item)

	if item.ID == "" {
		return false, base.NOT_FOUND_ERROR
	}

	r.DB.Delete(&item)

	return true, nil
}
