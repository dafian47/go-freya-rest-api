package handler

import (
	"github.com/labstack/echo"
	"net/http"

	base "github.com/dafian47/go-freya-rest-api/module"
	model "github.com/dafian47/go-freya-rest-api/module/event"
	"github.com/dafian47/go-freya-rest-api/module/event/repository"
)

type EventHandler interface {
	GetEventAll(c echo.Context) error
	GetEvent(c echo.Context) error
	CreateEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
	DeleteEvent(c echo.Context) error
}

type eventHandler struct {
	r repository.EventRepository
}

func NewEventHandler(e *echo.Echo, eventRepo repository.EventRepository) {

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
		return c.JSON(http.StatusNotFound, &base.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &base.Response{
		Status:  http.StatusOK,
		Message: base.SUCCESS_GET_DATA_ALL,
		Data:    itemList,
	})
}

func (h *eventHandler) GetEvent(c echo.Context) error {

	eventID := c.Param("id")

	item, err := h.r.GetEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    item,
		})
	}

	return c.JSON(http.StatusOK, &base.Response{
		Status:  http.StatusOK,
		Message: base.SUCCESS_GET_DATA,
		Data:    item,
	})
}

func (h *eventHandler) CreateEvent(c echo.Context) error {

	var item model.Event

	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	resultItem, err := h.r.CreateEvent(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &base.Response{
		Status:  http.StatusCreated,
		Message: base.SUCCESS_CREATE_DATA,
		Data:    resultItem,
	})
}

func (h *eventHandler) UpdateEvent(c echo.Context) error {

	var item model.Event

	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	resultItem, err := h.r.UpdateEvent(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &base.Response{
		Status:  http.StatusCreated,
		Message: base.SUCCESS_UPDATE_DATA,
		Data:    resultItem,
	})
}

func (h *eventHandler) DeleteEvent(c echo.Context) error {

	eventID := c.Param("id")

	_, err := h.r.GetEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	_, err = h.r.DeleteEvent(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &base.Response{
		Status:  http.StatusOK,
		Message: base.SUCCESS_DELETE_DATA,
		Data:    nil,
	})
}
