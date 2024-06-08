-- +goose Up

DROP VIEW diary_daily_summary;
-- +goose Down
CREATE OR REPLACE VIEW diary_daily_summary AS
SELECT
    date AS diary_date,
    SUM(food_moisture) AS total_food_moisture,
    SUM(salt) AS total_salt,
    SUM(total_cal) AS total_cal,
    AVG(score) AS average_score
FROM
    diary_menu
GROUP BY
    date;


