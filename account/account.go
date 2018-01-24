package account

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account domain struct
type Account struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
