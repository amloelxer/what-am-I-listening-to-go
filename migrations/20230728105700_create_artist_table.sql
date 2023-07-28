-- +goose Up
-- +goose StatementBegin
CREATE TABLE "artist" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "name" text, "biography" text, "createdAt" TIMESTAMP NOT NULL DEFAULT now(), "updatedAt" TIMESTAMP NOT NULL DEFAULT now(), PRIMARY KEY ("id"));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "artist"
-- +goose StatementEnd
