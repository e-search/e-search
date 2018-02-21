package repository_test

import (
	"regexp"
	"testing"

	"github.com/e-search/e-search/account"
	"github.com/e-search/e-search/account/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUpdate(t *testing.T) {
	u := &account.Account{
		ID:       uuid.NewV4(),
		Email:    "example@example.com",
		Password: "asdf0987",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "UPDATE accounts SET password = $1, email = $2, updated_at = now() WHERE id = $3"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(u.Password, u.Email, u.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	a := repository.NewAccountRepository(db)
	err = a.Update(u)
	assert.NoError(t, err)
	return
}
