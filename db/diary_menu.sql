-- name: CreateDiaryMenu :one
INSERT INTO diary_menu (diary_day_id, img, summary, total_cal,  menu_time, date, food_moisture, salt, score, carbon, fat, protein)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;
