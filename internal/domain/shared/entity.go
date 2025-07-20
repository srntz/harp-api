package shared

import "github.com/jackc/pgx/v5/pgtype"

type Entity struct {
	ID        string
	CreatedAt int
	UpdatedAt int
}

type EntitySQLC struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (e *Entity) ToSQLC() (*EntitySQLC, error) {
	var id pgtype.UUID
	if err := id.Scan(e.ID); err != nil {
		return nil, err
	}

	var createdAt pgtype.Timestamp
	if err := createdAt.Scan(e.CreatedAt); err != nil {
		return nil, err
	}

	var updatedAt pgtype.Timestamp
	if err := updatedAt.Scan(e.UpdatedAt); err != nil {
		return nil, err
	}

	return &EntitySQLC{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
