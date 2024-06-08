-- name: CreateDiaryDay :one
INSERT INTO diary_day (date)
VALUES ($1)
RETURNING *;

-- name: FindDiaryDayWithDate :one
SELECT * FROM diary_day WHERE date = $1;
