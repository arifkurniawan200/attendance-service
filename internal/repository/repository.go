package repository

import (
	"database/sql"
	"template/internal/model"
)

type UserRepository interface {
	RegisterUser(user model.MemberParam) error
	GetUserByEmail(email string) (model.Member, error)
	GetUserByID(userID int) (model.Member, error)
	GetUserExcludeMe(userID int) ([]model.Member, error)
	BeginTx() (*sql.Tx, error)
}

type GatheringRepository interface {
	CreateNewGatheringTx(tx *sql.Tx, g model.GatheringParam) (int64, error)
	AddAttendeeTx(tx *sql.Tx, a model.Attendee) error
	GetAttendee(gatheringID int) ([]model.Attendee, error)
	SendInvitation(tx *sql.Tx, data model.Invitation) error
	GetGathering(gatheringID int) (model.Gathering, error)
	UpdateInvitation(data model.Invitation) error
}
