-- name: GetArtistByID :one
SELECT * FROM artists WHERE id = $1;