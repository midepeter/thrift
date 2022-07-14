-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  id,
  email,
  password, 
  date_created
) VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: RemoveUser :one 
DELETE FROM users WHERE id = $1
RETURNING *;
