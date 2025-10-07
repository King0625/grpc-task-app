-- name: CreateUser :one
INSERT INTO users (email, password_hash, first_name, last_name, verify_token)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: VerifyUser :one
UPDATE users SET verified = true, verify_token = NULL
WHERE verify_token = $1
RETURNING *;