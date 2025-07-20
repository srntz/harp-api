package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/srntz/harp-api/internal/domain/artist"
	"github.com/srntz/harp-api/internal/infrastructure/db"
	"github.com/srntz/harp-api/internal/infrastructure/sqlc"
)

type ArtistRepository struct{}

func (r *ArtistRepository) GetByID(id string) (*artist.Artist, error) {
	pool, err := db.GetPool(db.DEFAULT_URL)
	if err != nil {
		return nil, err
	}

	q := sqlc.New(pool)

	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return nil, err
	}

	sqlcArtist, err := q.GetArtistByID(context.Background(), uuid)
	if err != nil {
		return nil, err
	}

	return artist.FromSQLC(&sqlcArtist), nil
}
