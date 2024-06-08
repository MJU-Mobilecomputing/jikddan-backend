-- +goose Up
DROP VIEW diary_daily_summary;
CREATE VIEW diary_daily_summary AS
SELECT
    date AS diary_date,
    SUM(food_moisture) AS total_food_moisture,
    SUM(salt) AS total_salt,
    SUM(total_cal) AS total_cal,
    SUM(carbon) AS total_carbon,
    SUM(fat) AS total_fat,
    SUM(protein) AS total_protein,
    AVG(score) AS average_score
FROM
    diary_menu
GROUP BY
    diary_date;

CREATE VIEW weekly_summary AS
SELECT 
    diary_date::date - EXTRACT(DOW FROM diary_date) * INTERVAL '1 DAY' AS start_date,
    diary_date::date + (6 - EXTRACT(DOW FROM diary_date)) * INTERVAL '1 DAY' AS end_date,
    EXTRACT(MONTH FROM diary_date) AS month,
    EXTRACT(WEEK FROM diary_date) - EXTRACT(WEEK FROM (DATE_TRUNC('MONTH', diary_date) + '1 DAY'::INTERVAL)::date - 1) AS week_of_month,
    SUM(total_food_moisture) AS weekly_food_moisture,
    SUM(total_salt) AS weekly_salt,
    SUM(total_carbon) AS weekly_carbon,
    SUM(total_fat) AS weekly_fat,
    SUM(total_protein) AS weekly_protein,
    AVG(average_score) AS weekly_average_score
FROM diary_daily_summary
GROUP BY diary_date - EXTRACT(DOW FROM diary_date) * INTERVAL '1 DAY', diary_date;

-- +goose Down
DROP VIEW weekly_summary;
DROP VIEW diary_daily_summary;


