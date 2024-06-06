-- +goose Up
CREATE TABLE "diary_day" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY("user_id") REFERENCES "user"(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "diary_day";

