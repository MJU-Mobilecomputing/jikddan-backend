-- name: CreateUser :one
INSERT INTO "user" (email, username, password)
VALUES ($1, $2, $3)
RETURNING id, email, username, created_at, updated_at;

-- name: FindUserById :one
SELECT id, email, username, created_at, updated_at FROM "user" WHERE id = $1;
