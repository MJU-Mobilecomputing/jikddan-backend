-- +goose Up
ALTER TABLE diary_menu
ADD COLUMN "food_moisture" int,
ADD COLUMN "salt" int,
ADD COLUMN "score" int,
DROP COLUMN "status";

-- +goose Down
ALTER TABLE diary_menu
DROP COLUMN "food_moisture",
DROP COLUMN "salt",
DROP COLUMN "score";

