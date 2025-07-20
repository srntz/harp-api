-- +goose Up
-- +goose StatementBegin
CREATE TYPE release_type as ENUM ('album', 'ep', 'single', 'compilation');

CREATE TABLE releases (
  id            UUID     PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at    TIMESTAMP     DEFAULT NOW(),
  updated_at    TIMESTAMP     DEFAULT NOW(),
  title         TEXT          NOT NULL,
  release_type  release_type  NOT NULL,
  image_url     TEXT CHECK (image_url SIMILAR TO 'https?://%'),
  artist_id     UUID     REFERENCES artists(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE releases;
DROP TYPE release_type;
-- +goose StatementEnd
