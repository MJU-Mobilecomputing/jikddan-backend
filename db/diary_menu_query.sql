-- name: CreateDiaryMenu :one
INSERT INTO diary_menu (img, summary, total_cal, status, menu_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
