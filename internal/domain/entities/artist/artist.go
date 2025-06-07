package artist

import "github.com/srntz/harp-api/internal/domain/shared"

type Artist struct {
	shared.Entity
	shared.Imagable
	name string
}
