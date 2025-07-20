package artist

import (
	"github.com/srntz/harp-api/internal/domain/shared"
	"github.com/srntz/harp-api/internal/infrastructure/sqlc"
)

type Artist struct {
	shared.Entity
	shared.ImageReferencer
	Name string
	Bio  string
}

func FromSQLC(entity *sqlc.Artist) *Artist {
	return &Artist{
		Entity: shared.Entity{
			ID:        entity.ID.String(),
			CreatedAt: entity.CreatedAt.Time.Nanosecond(),
			UpdatedAt: entity.UpdatedAt.Time.Nanosecond(),
		},
		ImageReferencer: shared.ImageReferencer{
			ImageURL: entity.ImageUrl.String,
		},
		Name: entity.Name,
		Bio:  entity.Bio.String,
	}
}

func (a *Artist) ToSQLC() (*sqlc.Artist, error) {
	entity, err := a.Entity.ToSQLC()
	if err != nil {
		return nil, err
	}

	imageURL, err := a.ImageReferencer.ToSQLC()
	if err != nil {
		return nil, err
	}

	return &sqlc.Artist{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		ImageUrl:  *imageURL,
	}, nil
}
