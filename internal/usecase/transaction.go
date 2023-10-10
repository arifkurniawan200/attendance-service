package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"template/internal/model"
	"template/internal/repository"
)

type GatheringHandler struct {
	g repository.GatheringRepository
	u repository.UserRepository
}

func NewTransactionsUsecase(g repository.GatheringRepository, u repository.UserRepository) GatheringUcase {
	return &GatheringHandler{g, u}
}

func (t GatheringHandler) CreateNewGathering(ctx *gin.Context, g model.GatheringParam) error {
	valid := model.CheckGatheringStatus(g.Type)
	if !valid {
		return fmt.Errorf("gathering type not valid,value should online/offline")
	}
	tx, err := t.u.BeginTx()
	if err != nil {
		return err
	}

	gatheringID, err := t.g.CreateNewGatheringTx(tx, g)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, userId := range g.AttendeeIDs {
		err = t.g.AddAttendeeTx(tx, model.Attendee{
			GatheringID: int(gatheringID),
			MemberID:    userId,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
