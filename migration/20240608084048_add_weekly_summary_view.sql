-- +goose Up
ALTER TABLE diary_menu
ADD COLUMN "carbon" int,
ADD COLUMN "fat" int,
ADD COLUMN "protein" int;

-- +goose Down
ALTER TABLE diary_menu
DROP COLUMN "carbon",
DROP COLUMN "fat",
DROP COLUMN "protein";


