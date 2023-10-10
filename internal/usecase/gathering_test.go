package usecase

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"template/internal/model"
	"testing"
)

// Mock GatheringRepository
type MockGatheringRepository struct{}

func (r *MockGatheringRepository) GetGatheringInfo(gatheringID int) (interface{}, error) {
	return "ok", nil
}

func (r *MockGatheringRepository) GetGathering(gatheringID int) (model.Gathering, error) {
	return model.Gathering{}, nil
}

func (r *MockGatheringRepository) UpdateInvitation(invitation model.Invitation) error {
	return nil
}

func (r *MockGatheringRepository) SendInvitation(tx *sql.Tx, invitation model.Invitation) error {
	return nil
}

func (r *MockGatheringRepository) GetAttendee(gatheringID int) ([]model.Attendee, error) {
	return []model.Attendee{}, nil
}

func (r *MockGatheringRepository) CreateNewGatheringTx(tx *sql.Tx, g model.GatheringParam) (int64, error) {
	return 1, nil
}

func (r *MockGatheringRepository) AddAttendeeTx(tx *sql.Tx, attendee model.Attendee) error {
	return nil
}

// Unit test untuk GetGatheringInfo
func TestGetGatheringInfo(t *testing.T) {
	ctx := &gin.Context{}

	gatheringRepo := &MockGatheringRepository{}

	userRepo := &MockUserRepository{}

	gatheringUcase := NewGatheringUsecase(gatheringRepo, userRepo)

	gatheringID := 1
	info, err := gatheringUcase.GetGatheringInfo(ctx, gatheringID)

	assert.NoError(t, err)

	assert.NotNil(t, info)
}

// Unit test untuk ApproveInvitation
func TestApproveInvitation(t *testing.T) {
	ctx := &gin.Context{}
	gatheringRepo := &MockGatheringRepository{}
	userRepo := &MockUserRepository{}
	gatheringUcase := NewGatheringUsecase(gatheringRepo, userRepo)
	invitation := model.Invitation{
		MemberID:    1,
		GatheringID: 1,
		Status:      model.InvitationStatusApprove,
	}

	err := gatheringUcase.ApproveInvitation(ctx, invitation)

	assert.NoError(t, err)
}

// Unit test untuk RejectInvitation
func TestRejectInvitation(t *testing.T) {
	ctx := &gin.Context{}

	gatheringRepo := &MockGatheringRepository{}

	userRepo := &MockUserRepository{}

	gatheringUcase := NewGatheringUsecase(gatheringRepo, userRepo)

	invitation := model.Invitation{
		MemberID:    1,
		GatheringID: 1,
		Status:      model.InvitationStatusReject,
	}

	err := gatheringUcase.RejectInvitation(ctx, invitation)

	assert.NoError(t, err)
}

// Unit test untuk SendInvitation
func TestSendInvitation(t *testing.T) {
	ctx := &gin.Context{}

	gatheringRepo := &MockGatheringRepository{}

	userRepo := &MockUserRepository{}

	gatheringUcase := NewGatheringUsecase(gatheringRepo, userRepo)

	userID := 1
	gatheringID := 1
	err := gatheringUcase.SendInvitation(ctx, userID, gatheringID)

	assert.Error(t, err)
}
