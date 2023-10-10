package model

type InvitationStatus string

const (
	InvitationStatusSent    InvitationStatus = "Sent"
	InvitationStatusApprove InvitationStatus = "Approve"
	InvitationStatusReject  InvitationStatus = "Reject"
)

type Invitation struct {
	ID          int              `json:"id"`
	MemberID    int              `json:"member_id"`
	GatheringID int              `json:"gathering_id"`
	Status      InvitationStatus `json:"status"`
}
