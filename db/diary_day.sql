-- name: CreateDiaryDay :one
INSERT INTO diary_day (date)
VALUES ($1)
RETURNING *;

-- name: FindDiaryDayWithDate :one
SELECT * FROM diary_day WHERE date = $1;

-- name: FindDiaryDayWithMenus :one
SELECT * FROM diary_day_view WHERE diary_date = $1;

-- name: FindDailySummaryWithDate :one
SELECT * FROM diary_daily_summary WHERE diary_date = $1;
