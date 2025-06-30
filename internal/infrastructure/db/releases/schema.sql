CREATE TYPE release_type as ENUM ('album', 'ep', 'single', 'compilation');

CREATE TABLE releases (
  id            BIGSERIAL     PRIMARY KEY,
  created_at    TIMESTAMP     DEFAULT NOW(),
  updated_at    TIMESTAMP     DEFAULT NOW(),
  title         TEXT          NOT NULL,
  release_type  release_type  NOT NULL,
  image_url     TEXT CHECK (image_url SIMILAR TO 'https?://%'),
  artist_id     UUID     REFERENCES artists(id)
);

