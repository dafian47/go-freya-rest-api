package handler

import (
	"github.com/dafian47/go-freya-rest-api/module/event/repository/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEventAll(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/event/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(mocks.EventRepository)
	//mockRepo.On("GetEventAll").Return(&mockRepo, nil)

	h := eventHandler{
		r: mockRepo,
	}

	h.GetEventAll(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetEvent(t *testing.T) {

}

func TestCreateEvent(t *testing.T) {

}

func TestUpdateEvent(t *testing.T) {

}

func TestDeleteEvent(t *testing.T) {

}
