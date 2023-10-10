package repository

const (
	insertNewMembers = `INSERT INTO members (first_name, last_name, email, password) VALUES
    (?, ?, ?, ?)`
	getMembersByEmail = `
			SELECT id, first_name, last_name, email,password, created_at, updated_at,deleted_at
	FROM members WHERE email = ?
    `
)
