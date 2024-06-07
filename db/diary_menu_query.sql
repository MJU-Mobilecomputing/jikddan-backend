-- name: CreateDiaryMenu :one
INSERT INTO diary_menu (img, summary, total_cal, status, menu_time, diary_day_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
