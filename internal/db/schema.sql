CREATE TYPE release_type as ENUM ('album', 'ep', 'single', 'compilation');

CREATE TABLE artists (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  name          TEXT        NOT NULL,
  description   TEXT,
  image_url     VARCHAR(255)
);

CREATE TABLE releases (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  title         TEXT        NOT NULL,
  playtime      INTEGER     NOT NULL,
  image_url     VARCHAR(255),
  artist_id     BIGSERIAL   REFERENCES artists(id)
);

CREATE TABLE tracks (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  title         TEXT        NOT NULL,
  playtime_sec  INTEGER     NOT NULL,
  release_id    BIGSERIAL   REFERENCES releases(id)
);

CREATE TABLE genre (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  name          TEXT        UNIQUE NOT NULL
);

CREATE TABLE subgenre (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  name          TEXT        UNIQUE NOT NULL, 
  genre_id      BIGSERIAL   REFERENCES genre(id)
);

CREATE TABLE users (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  username      VARCHAR(64) UNIQUE NOT NULL,
  bio           TEXT
);

CREATE TABLE reviews (
  id            BIGSERIAL   PRIMARY KEY,
  created_at    DATETIME    DEFAULT NOW(),
  updated_at    DATETIME    DEFAULT NOW(),
  title         VARCHAR(255)  NOT NULL,
  body          TEXT
);
