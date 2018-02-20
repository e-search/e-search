package usecase_test

import (
	"testing"

	"github.com/e-search/e-search/account"
	"github.com/e-search/e-search/account/repository/mocks"
	"github.com/e-search/e-search/account/usecase"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	mockAccountRepo := new(mocks.AccountRepository)
	mockAccount := account.Account{
		ID:       uuid.NewV4(),
		Email:    "example@example.com",
		Password: "asdf0987",
	}

	mockAccountRepo.On("Update", mock.AnythingOfType("*account.Account")).Return(nil)
	defer mockAccountRepo.AssertCalled(t, "Update", mock.AnythingOfType("*account.Account"))

	u := usecase.NewAccountUsecase(mockAccountRepo)
	err := u.Update(&mockAccount)
	assert.NoError(t, err)
}
