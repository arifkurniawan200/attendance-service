package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"template/internal/model"
	"template/internal/repository"
	"template/internal/utils"
)

type GatheringHandler struct {
	g repository.GatheringRepository
	u repository.UserRepository
}

func NewGatheringUsecase(g repository.GatheringRepository, u repository.UserRepository) GatheringUcase {
	return &GatheringHandler{g, u}
}

func (t GatheringHandler) ApproveInvitation(ctx *gin.Context, data model.Invitation) error {
	gathering, err := t.g.GetGathering(data.GatheringID)
	if err != nil {
		return err
	}

	creator, err := t.u.GetUserByID(gathering.Creator)
	if err != nil {
		return err
	}

	err = t.g.UpdateInvitation(model.Invitation{
		Status:      model.InvitationStatusApprove,
		GatheringID: data.GatheringID,
		MemberID:    data.MemberID,
	})
	if err != nil {
		return err
	}

	go func(notification utils.Notification) {
		err = utils.SendNotification(notification)
		if err != nil {
			return
		}
	}(utils.Notification{
		Type:    utils.NotificationTypeEmail,
		Subject: "Confirmation Join Gathering",
		Target:  creator.Email,
		Body:    fmt.Sprintf("Someone have confirmation join %s, check your gathering information to show details", gathering.Name),
	})

	return err
}

func (t GatheringHandler) RejectInvitation(ctx *gin.Context, data model.Invitation) error {
	gathering, err := t.g.GetGathering(data.GatheringID)
	if err != nil {
		return err
	}

	creator, err := t.u.GetUserByID(gathering.Creator)
	if err != nil {
		return err
	}

	err = t.g.UpdateInvitation(model.Invitation{
		Status:      model.InvitationStatusReject,
		GatheringID: data.GatheringID,
		MemberID:    data.MemberID,
	})
	if err != nil {
		return err
	}

	go func(notification utils.Notification) {
		err = utils.SendNotification(notification)
		if err != nil {
			return
		}
	}(utils.Notification{
		Type:    utils.NotificationTypeEmail,
		Subject: "Reject Gathering",
		Target:  creator.Email,
		Body:    fmt.Sprintf("Someone have confirmation and reject your invitation gathering %s, check your gathering information to show details", gathering.Name),
	})

	return err
}

func (t GatheringHandler) SendInvitation(ctx *gin.Context, userID, gatheringID int) error {
	gathering, err := t.g.GetGathering(gatheringID)
	if err != nil {
		return err
	}
	if gathering.Creator != userID {
		return fmt.Errorf("your not creatof of the gathering, only creator allowed to send invitation")
	}

	attendes, err := t.g.GetAttendee(gathering.ID)
	if err != nil {
		return err
	}

	tx, err := t.u.BeginTx()
	if err != nil {
		return err
	}

	for _, att := range attendes {
		err = t.g.SendInvitation(tx, model.Invitation{
			MemberID:    att.MemberID,
			GatheringID: gathering.ID,
			Status:      model.InvitationStatusSent,
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

	go func(gathering model.Gathering, attendees []model.Attendee) {
		for _, att := range attendees {
			user, err := t.u.GetUserByID(att.MemberID)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			err = utils.SendNotification(utils.Notification{
				Type:    utils.NotificationTypeEmail,
				Subject: "You receive new invitation",
				Body:    fmt.Sprintf("you have received invitation gathering %s, the event will be held on %s %v", gathering.Name, gathering.Location, gathering.ScheduleAt),
				Target:  user.Email,
			})
			if err != nil {
				log.Error(err.Error())
				continue
			}
		}
	}(gathering, attendes)

	return err
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
