package mocks

import (
	"github.com/e-search/e-search/account"
	"github.com/stretchr/testify/mock"
)

// AccountRepository mock account repository
type AccountRepository struct {
	mock.Mock
}

// Update provide a mock function
func (m *AccountRepository) Update(u *account.Account) error {
	ret := m.Called(u)

	// var r0 *account.Account
	// rf, ok := ret.Get(0).(func(*account.Account) *account.Account)
	// if ok {
	// 	r0 = rf(u)
	// } else {
	// 	if ret.Get(0) != nil {
	// 		r0 = ret.Get(0).(*account.Account)
	// 	}
	// }

	var r1 error
	rf, ok := ret.Get(1).(func(*account.Account) error)
	if ok {
		r1 = rf(u)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Error(1)
		}
	}
	return r1
}
