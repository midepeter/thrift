-- name: GetName :one
SELECT * FROM users
WHERE id = $1 LIMIT $1;

-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  password,
  phone_number,
  created_at,
  updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: RemoveUser :exec
DELETE FROM users WHERE id = $1
RETURNING *;
