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
}
