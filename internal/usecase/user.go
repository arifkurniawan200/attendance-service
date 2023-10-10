package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"template/internal/model"
	"template/internal/repository"
	"template/internal/utils"
)

type UserHandler struct {
	u repository.UserRepository
	t repository.TransactionRepository
}

const (
	secret = "abc&1*~#^2^#s0^=)^^7%b34"
)

func NewUserUsecase(u repository.UserRepository, t repository.TransactionRepository) UserUcase {
	return &UserHandler{u, t}
}

func (u UserHandler) GetUserInfoByEmail(ctx *gin.Context, email string) (model.Member, error) {
	var (
		err error
	)
	userInfo, err := u.u.GetUserByEmail(email)
	if err != nil {
		return model.Member{}, err
	}
	return userInfo, err
}

func (u UserHandler) RegisterCustomer(ctx *gin.Context, c model.MemberParam) error {
	var (
		err error
		g   errgroup.Group
	)

	g.Go(func() error {
		// hash password
		c.Password, err = utils.HashPassword(c.Password)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return err
	}

	err = u.u.RegisterUser(c)
	if err != nil {
		log.Errorf("[usecase][RegisterCustomer] error when RegisterUser: %s", err.Error())
		return err
	}

	return nil
}
