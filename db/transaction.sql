-- name: GetBalance :one
SELECT * FROM balances
WHERE user_id = $1 LIMIT 1;

-- name: CreateTransaction :one
INSERT INTO transactions (
  transaction_id,
  user_id,
  currency_id,
  transaction_amount,
  transaction_date
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;



