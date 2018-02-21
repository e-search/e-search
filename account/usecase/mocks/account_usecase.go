package mocks

import (
	"github.com/e-search/e-search/account"
	"github.com/stretchr/testify/mock"
)

// AccountUsecase mock account usecase
type AccountUsecase struct {
	mock.Mock
}

// Create provide a mock function
func (m *AccountUsecase) Create(u *account.Account) (*account.Account, error) {
	ret := m.Called(u)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(*account.Account) *account.Account); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*account.Account) error); ok {
		r1 = rf(u)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Error(1)
		}
	}
	return r0, r1
}

// Update provide a mock function
func (m *AccountUsecase) Update(u *account.Account) error {
	ret := m.Called(u)

	var r0 error
	rf, ok := ret.Get(0).(func(*account.Account) error)
	if ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(0)
		}
	}
	return r0
}
