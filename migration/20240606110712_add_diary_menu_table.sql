-- +goose Up
CREATE TYPE status AS ENUM ('pending','complete');
CREATE TYPE menu_time AS ENUM ('breakfast','lunch','dinner','snack');
CREATE TABLE "diary_menu" (
  "id" bigserial PRIMARY KEY,
  "diary_day_id" bigint NOT NULL,
  "img" varchar NOT NULL,
  "summary" varchar NOT NULL,
  "total_cal" int NOT NULL,
  "status" status,
  "menu_time" menu_time,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY("diary_day_id") REFERENCES "diary_day"(id) ON DELETE CASCADE
);

-- +goose Down
DROP TYPE status;
DROP TABLE "diary_menu";

