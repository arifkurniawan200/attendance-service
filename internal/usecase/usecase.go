package usecase

import (
	"github.com/gin-gonic/gin"
	"template/internal/model"
)

type UserUcase interface {
	RegisterCustomer(ctx *gin.Context, customer model.UserParam) error
	GetUserInfoByEmail(ctx *gin.Context, email string) (model.User, error)
}

type TransactionUcase interface {
}
