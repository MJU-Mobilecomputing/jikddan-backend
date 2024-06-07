-- name: CreateDiaryDay :one
INSERT INTO "diary_day" (user_id, "date")
VALUES ($1, $2)
RETURNING *;

-- name: FindOneDiaryWithMenu :one
SELECT * FROM diary_day_view WHERE diary_day_id = $1 AND EXTRACT(MONTH FROM diary_date) = sqlc.arg(month)::int AND EXTRACT(DAY FROM diary_date) = sqlc.arg(day)::int;
