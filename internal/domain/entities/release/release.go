package release

import "github.com/srntz/harp-api/internal/domain/shared"

type Release struct {
	shared.Entity
	shared.Imagable
	title        string
	release_type ReleaseType
	artist_id    shared.EntityIndentifier
}
