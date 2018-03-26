package handler

import (
	model "github.com/dafian47/go-freya-rest-api/module/event"
	repo "github.com/dafian47/go-freya-rest-api/module/event/repository"
	"github.com/labstack/echo"
	"net/http"
)

type EventHandler interface {
	GetEventAll(c echo.Context) error
	GetEvent(c echo.Context) error
	CreateEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
	DeleteEvent(c echo.Context) error
}

type eventHandler struct {
	r repo.EventRepository
}

func NewEventHandler(e *echo.Echo, eventRepo repo.EventRepository) {

	handler := &eventHandler{r: eventRepo}

	e.GET("/event/", handler.GetEventAll)
	e.GET("/event/:id", handler.GetEvent)
	e.POST("/event/", handler.CreateEvent)
	e.PUT("/event/:id", handler.UpdateEvent)
	e.DELETE("/event/:id", handler.DeleteEvent)
}

func (h *eventHandler) GetEventAll(c echo.Context) error {

	itemList, err := h.r.GetEventAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, itemList)
}

func (h *eventHandler) GetEvent(c echo.Context) error {

	eventID := c.Param("id")

	item, err := h.r.GetEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *eventHandler) CreateEvent(c echo.Context) error {

	var item model.Event

	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resultItem, err := h.r.CreateEvent(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resultItem)
}

func (h *eventHandler) UpdateEvent(c echo.Context) error {

	var item model.Event

	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resultItem, err := h.r.UpdateEvent(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resultItem)
}

func (h *eventHandler) DeleteEvent(c echo.Context) error {

	eventID := c.Param("id")

	_, err := h.r.GetEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	_, err = h.r.DeleteEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Success delete event")
}
