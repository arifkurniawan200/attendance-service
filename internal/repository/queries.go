package repository

const (
	insertNewMembers = `INSERT INTO members (first_name, last_name, email, password) VALUES
    (?, ?, ?, ?)`
	baseGetMember = `
			SELECT id, first_name, last_name, email,password, created_at, updated_at,deleted_at
	FROM members %s
    `
	createNewGathering = `
INSERT INTO gatherings (creator, type, schedule_at, name, location)
VALUES (?, ?, ?, ?, ?);
`
	addGatheringAttendee = `INSERT INTO attendees (member_id, gathering_id)
VALUES (?, ?); `
)
