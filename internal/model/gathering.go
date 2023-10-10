package model

import (
	"time"
)

type GatheringStatus string

const (
	GatheringStatusOnline  GatheringStatus = "online"
	GatheringStatusOffline GatheringStatus = "offline"
)

func CheckGatheringStatus(status GatheringStatus) bool {
	checkGatheringStatus := map[GatheringStatus]GatheringStatus{
		"online":  GatheringStatusOnline,
		"offline": GatheringStatusOffline,
	}
	_, found := checkGatheringStatus[status]
	if !found {
		return false
	}
	return true
}

type Gathering struct {
	ID         int             `json:"id"`
	Creator    int             `json:"creator"`
	Type       GatheringStatus `json:"type"`
	ScheduleAt time.Time       `json:"schedule_at"`
	Name       string          `json:"name"`
	Location   string          `json:"location"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

type GatheringParam struct {
	Creator     int64           `json:"-" `
	Type        GatheringStatus `json:"type" validate:"required"`
	ScheduleAt  time.Time       `json:"schedule_at" validate:"required"`
	Name        string          `json:"name" validate:"required"`
	Location    string          `json:"location" validate:"required"`
	AttendeeIDs []int           `json:"attendee_ids" validate:"required"`
}

type Attendee struct {
	MemberID    int
	GatheringID int
}
