package usecase

import (
	"github.com/gin-gonic/gin"
	"template/internal/model"
)

type UserUcase interface {
	RegisterCustomer(ctx *gin.Context, customer model.MemberParam) error
	GetUserInfoByEmail(ctx *gin.Context, email string) (model.Member, error)
}

type GatheringUcase interface {
	CreateNewGathering(ctx *gin.Context, g model.GatheringParam) error
	SendInvitation(ctx *gin.Context, userID, gatheringID int) error
	ApproveInvitation(ctx *gin.Context, data model.Invitation) error
	RejectInvitation(ctx *gin.Context, data model.Invitation) error
}
