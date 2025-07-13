-- +goose Up
-- +goose StatementBegin
CREATE TABLE artists (
  id            UUID   PRIMARY KEY,
  created_at    TIMESTAMP    DEFAULT NOW(),
  updated_at    TIMESTAMP    DEFAULT NOW(),
  name          TEXT        NOT NULL,
  bio           TEXT,
  image_url     TEXT CHECK (image_url SIMILAR TO 'https?://%')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE artists;
-- +goose StatementEnd
