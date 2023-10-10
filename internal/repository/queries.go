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

	getAttendee = `SELECT member_id,gathering_id FROM attendees WHERE gathering_id = ?;`

	createNewInvitation = `INSERT INTO invitations (member_id, gathering_id, status) VALUES (?,?,?);`

	getGatheringByUserID = `SELECT id, creator, type, schedule_at, name, location, created_at, updated_at, deleted_at FROM gatherings
WHERE id = ?
`
	queryUpdateInvitation = `UPDATE invitations
							SET status = ?
							WHERE gathering_id = ? AND member_id = ?;`

	queryGetGatheringInfo = `
		SELECT
		  'approve' AS status,
		  IFNULL(JSON_ARRAYAGG(u.email), JSON_ARRAY()) AS emails
		FROM
		  invitations i
		INNER JOIN
		  users u ON i.member_id = u.id
		WHERE
		  i.status = 'Approve'
		  AND i.gathering_id = ?
		UNION
		SELECT
		  'reject' AS status,
		  IFNULL(JSON_ARRAYAGG(u.email), JSON_ARRAY()) AS emails
		FROM
		  invitations i
		INNER JOIN
		  users u ON i.member_id = u.id
		WHERE
		  i.status = 'Reject'
		  AND i.gathering_id = ?
		UNION
		SELECT
		  'sent' AS status,
		  IFNULL(JSON_ARRAYAGG(u.email), JSON_ARRAY()) AS emails
		FROM
		  invitations i
		INNER JOIN
		  users u ON i.member_id = u.id
		WHERE
		  i.status = 'Sent'
		  AND i.gathering_id = ?`
)
