-- name: FindWeeklySummary :one
SELECT month,week_of_month,weekly_food_moisture,weekly_salt,weekly_carbon,weekly_fat,weekly_protein,weekly_average_score FROM weekly_summary WHERE month = sqlc.arg(month)::int AND week_of_month = sqlc.arg(week_num)::int;
