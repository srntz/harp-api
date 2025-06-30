-- name: GetArtists :many
SELECT * FROM artists
OFFSET $1 LIMIT $2;

-- name: GetArtistsCount :one
SELECT COUNT(*) FROM artists;

-- name: GetArtistByID :one
SELECT * FROM artists
WHERE id = $1;
