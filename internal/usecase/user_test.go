package usecase

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"template/internal/model"
	"testing"
)

// Mock UserRepository
type MockUserRepository struct{}

func (r *MockUserRepository) GetUserByID(userID int) (model.Member, error) {
	return model.Member{}, nil
}

func (r *MockUserRepository) BeginTx() (*sql.Tx, error) {
	return &sql.Tx{}, nil
}

func (r *MockUserRepository) GetUserExcludeMe(userID int) ([]model.Member, error) {
	return []model.Member{}, nil
}

func (r *MockUserRepository) GetUserByEmail(email string) (model.Member, error) {
	return model.Member{}, nil
}

func (r *MockUserRepository) RegisterUser(c model.MemberParam) error {
	return nil
}

// Unit test untuk GetUserFriends
func TestGetUserFriends(t *testing.T) {
	ctx := &gin.Context{}

	userRepo := &MockUserRepository{}

	userUcase := NewUserUsecase(userRepo, nil)

	userID := 1
	friends, err := userUcase.GetUserFriends(ctx, userID)

	assert.NoError(t, err)

	assert.IsType(t, []model.Member{}, friends)
}

// Unit test untuk GetUserInfoByEmail
func TestGetUserInfoByEmail(t *testing.T) {
	ctx := &gin.Context{}

	userRepo := &MockUserRepository{}
	userUcase := NewUserUsecase(userRepo, nil)

	email := "user@example.com"
	userInfo, err := userUcase.GetUserInfoByEmail(ctx, email)
	assert.NoError(t, err)
	assert.IsType(t, model.Member{}, userInfo)
}

// Unit test untuk RegisterCustomer
func TestRegisterCustomer(t *testing.T) {
	ctx := &gin.Context{}
	userRepo := &MockUserRepository{}
	userUcase := NewUserUsecase(userRepo, nil)
	customer := model.MemberParam{
		FirstName: "Arif",
	}

	err := userUcase.RegisterCustomer(ctx, customer)

	assert.NoError(t, err)
}
