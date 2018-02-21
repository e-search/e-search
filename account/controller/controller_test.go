package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/e-search/e-search/account"
	"github.com/e-search/e-search/account/controller"
	"github.com/e-search/e-search/account/usecase/mocks"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	mockAccount := account.Account{
		ID:       uuid.NewV4(),
		Email:    "example@example.com",
		Password: "asdf0987",
	}
	mockUcase := new(mocks.AccountUsecase)
	mockUcase.On("Update", mock.AnythingOfType("*account.Account")).Return(nil)

	j, err := json.Marshal(&mockAccount)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.PATCH, "/account", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := controller.AccountController{
		AccountUsecase: mockUcase,
	}

	handler.Update(c)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUcase.AssertExpectations(t)
}
