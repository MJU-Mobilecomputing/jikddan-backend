-- +goose Up
CREATE VIEW diary_day_view AS
SELECT
    d.id AS diary_day_id,
    d.date AS diary_date,
    JSON_AGG(DISTINCT m.*) AS diary_menus
FROM
    diary_day d
LEFT JOIN
    diary_menu m ON d.id = m.diary_day_id
GROUP BY d.id
ORDER BY
    diary_date, diary_day_id;
-- +goose Down
DROP VIEW diary_day_view;
