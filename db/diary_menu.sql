-- name: CreateDiaryMenu :one
INSERT INTO diary_menu (diary_day_id, img, summary, total_cal, status, menu_time, date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
