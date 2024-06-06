-- name: CreateDiaryDay :one
INSERT INTO "diary_day" (user_id)
VALUES ($1)
RETURNING *;

-- name: FindOneDiaryWithMenu :one
SELECT * FROM diary_day_view WHERE diary_day_id = $1 AND EXTRACT(MONTH FROM diary_data) = sqlc.arg(month)::int;

