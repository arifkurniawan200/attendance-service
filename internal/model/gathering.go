package model

import "time"

type GatheringStatus string

const (
	GatheringStatusOnline  GatheringStatus = "online"
	GatheringStatusOffline GatheringStatus = "offline"
)

type Gathering struct {
	ID         int             `json:"id"`
	Creator    string          `json:"creator"`
	Type       GatheringStatus `json:"type"`
	ScheduleAt time.Time       `json:"schedule_at"`
	Name       string          `json:"name"`
	Location   string          `json:"location"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}
