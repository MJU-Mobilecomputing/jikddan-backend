-- name: CreateUser :one
INSERT INTO "user" (email, username, password, tall, weight)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, username, tall, weight, created_at, updated_at;

-- name: FindUserById :one
SELECT id, email, username, tall, weight, created_at, updated_at FROM "user" WHERE id = $1;
