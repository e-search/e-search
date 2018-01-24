package repository

import (
	"database/sql"

	"github.com/e-search/e-search/account"
)

// AccountRepository repository interface
type AccountRepository interface {
	Create(*account.Account) (*account.Account, error)
}

type accountRepository struct {
	Conn *sql.DB
}

// NewAccountRepository new accountrepository
func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

func (m *accountRepository) Create(u *account.Account) (*account.Account, error) {
	query := `insert into accounts (id, password, email, created_at, updated_at) values (uuid_generate_v4(), $1, $2, now(), now())`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(u.Password, u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
