-- name: CreateUser :one
INSERT INTO "user" (email, username, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindUserById :one
SELECT * FROM "user" WHERE id = $1;
