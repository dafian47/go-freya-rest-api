package handler

import (
	"github.com/dafian47/go-freya-rest-api/module/event/repository/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventHandler_GetEventAll(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/event/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(mocks.EventRepository)

	h := eventHandler{
		r: mockRepo,
	}

	h.GetEventAll(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestEventHandler_GetEvent(t *testing.T) {

	id := "02"

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/event/"+id, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/event/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	mockRepo := new(mocks.EventRepository)

	h := eventHandler{
		r: mockRepo,
	}

	h.GetEvent(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestEventHandler_CreateEvent(t *testing.T) {

}

func TestEventHandler_UpdateEvent(t *testing.T) {

}

func TestEventHandler_DeleteEvent(t *testing.T) {

}
