package usecase

import (
	"errors"

	"github.com/e-search/e-search/account"
	"github.com/e-search/e-search/account/repository"
	uuid "github.com/satori/go.uuid"
)

// AccountUsecase usecase interface
type AccountUsecase interface {
	Create(*account.Account) (*account.Account, error)
	Update(*account.Account) error
}

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

// NewAccountUsecase mount account repository
func NewAccountUsecase(account repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: account,
	}
}

func (a *accountUsecase) Create(account *account.Account) (*account.Account, error) {
	err := a.validateCreate(account)
	if err != nil {
		return nil, err
	}

	res, err := a.accountRepo.Create(account)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update update account
func (a *accountUsecase) Update(account *account.Account) error {
	err := a.validateUpdate(account)
	if err != nil {
		return err
	}

	err = a.accountRepo.Update(account)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountUsecase) validateCreate(account *account.Account) error {
	if account.Email == "" || account.Password == "" {
		return errors.New("email or password is required")
	}
	return nil
}

func (a *accountUsecase) validateUpdate(account *account.Account) error {
	if account.ID == uuid.FromStringOrNil("") || account.Email == "" || account.Password == "" {
		return errors.New("id, email or passowrd is required")
	}
	return nil
}
