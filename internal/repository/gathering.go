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

func (g GatheringHandler) UpdateInvitation(data model.Invitation) error {
	_, err := g.db.Exec(queryUpdateInvitation, data.Status, data.GatheringID, data.MemberID)
	if err != nil {
		return err
	}
	return err
}

func (g GatheringHandler) GetGathering(gatheringID int) (model.Gathering, error) {
	var (
		data model.Gathering
		err  error
	)
	rows, err := g.db.Query(getGatheringByUserID, gatheringID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.ID, &data.Creator, &data.Type, &data.ScheduleAt, &data.Name, &data.Location, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt); err != nil {
			return data, err
		}
	}

	if err = rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (g GatheringHandler) SendInvitation(tx *sql.Tx, data model.Invitation) error {
	_, err := tx.Exec(createNewInvitation, data.MemberID, data.GatheringID, data.Status)
	if err != nil {
		return err
	}
	return err
}

func (g GatheringHandler) GetAttendee(gatheringID int) ([]model.Attendee, error) {
	var (
		datas []model.Attendee
		err   error
	)
	rows, err := g.db.Query(getAttendee, gatheringID)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.Attendee
		if err = rows.Scan(&data.MemberID, &data.GatheringID); err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		return datas, err
	}
	return datas, err
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
