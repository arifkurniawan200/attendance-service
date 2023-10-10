package repository

import (
	"database/sql"
	"template/internal/model"
)

type UserRepository interface {
	RegisterUser(user model.MemberParam) error
	GetUserByEmail(email string) (model.Member, error)
	BeginTx() (*sql.Tx, error)
}

type GatheringRepository interface {
}
