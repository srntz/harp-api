CREATE TYPE release_type as ENUM ('album', 'ep', 'single', 'compilation');

CREATE TABLE artists (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    TIMESTAMP    DEFAULT NOW(),
  updated_at    TIMESTAMP    DEFAULT NOW(),
  name          TEXT        NOT NULL,
  image_url     VARCHAR(255)
);

CREATE TABLE releases (
  id            BIGSERIAL     PRIMARY KEY,
  created_at    TIMESTAMP     DEFAULT NOW(),
  updated_at    TIMESTAMP     DEFAULT NOW(),
  title         TEXT          NOT NULL,
  release_type  release_type  NOT NULL,
  image_url     VARCHAR(255),
  artist_id     BIGSERIAL     REFERENCES artists(id)
);

CREATE TABLE tracks (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    TIMESTAMP   DEFAULT NOW(),
  updated_at    TIMESTAMP   DEFAULT NOW(),
  title         TEXT        NOT NULL,
  playtime_sec  INTEGER     NOT NULL,
  release_id    BIGSERIAL   REFERENCES releases(id)
);

CREATE TABLE users (
  id            BIGSERIAL     PRIMARY KEY,
  created_at    TIMESTAMP     DEFAULT NOW(),
  updated_at    TIMESTAMP     DEFAULT NOW(),
  username      VARCHAR(64)   UNIQUE NOT NULL,
  bio           TEXT
);

CREATE TABLE reviews (
  id            BIGSERIAL     PRIMARY KEY,
  created_at    TIMESTAMP     DEFAULT NOW(),
  updated_at    TIMESTAMP     DEFAULT NOW(),
  header        VARCHAR(255)  NOT NULL,
  body          TEXT
  user_id       BIGSERIAL     REFERENCES users(id)
);

-- The following tables are currently not in use
-- and will need polishing before being rolled out.
CREATE TABLE genres (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    TIMESTAMP   DEFAULT NOW(),
  updated_at    TIMESTAMP   DEFAULT NOW(),
  name          TEXT        UNIQUE NOT NULL
);

CREATE TABLE subgenres (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    TIMESTAMP   DEFAULT NOW(),
  updated_at    TIMESTAMP   DEFAULT NOW(),
  name          TEXT        UNIQUE NOT NULL, 
  genre_id      BIGSERIAL   REFERENCES genre(id)
);

