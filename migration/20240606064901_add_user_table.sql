-- +goose Up
CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL UNIQUE,
  "username" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "tall" int NOT NULL,
  "weight" int NOT NULL,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE "user";

