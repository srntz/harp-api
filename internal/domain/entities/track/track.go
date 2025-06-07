package track

import "github.com/srntz/harp-api/internal/domain/shared"

type Track struct {
	shared.Entity
	title        string
	playtime_sec int
	release_id   shared.EntityIndentifier
}
