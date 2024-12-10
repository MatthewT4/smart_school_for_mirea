package storage

const (
	queryCreateUser = `
INSERT INTO users (id, email, password)
VALUES ($1, $2, $3)
RETURNING id, email, password
`

	queryGetUserByEmail = `
SELECT id, email, password
FROM users
WHERE email = $1
`

	queryGetUserByID = `
SELECT id, email, password
FROM users
WHERE id = $1
`
)
