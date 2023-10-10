package repository

import (
	"database/sql"
	"template/internal/model"
)

type GatheringHandler struct {
	db *sql.DB
}

func NewGatheringRepository(db *sql.DB) GatheringRepository {
	return &GatheringHandler{db}
}

func (g GatheringHandler) CreateNewGatheringTx(tx *sql.Tx, gth model.GatheringParam) (int64, error) {
	rows, err := tx.Exec(createNewGathering, gth.Creator, gth.Type, gth.ScheduleAt, gth.Name, gth.Location)
	if err != nil {
		return 0, err
	}
	return rows.LastInsertId()
}

func (g GatheringHandler) AddAttendeeTx(tx *sql.Tx, a model.Attendee) error {
	_, err := tx.Exec(addGatheringAttendee, a.MemberID, a.GatheringID)
	if err != nil {
		return err
	}
	return err
}
