package shared

import "github.com/jackc/pgx/v5/pgtype"

type ImageReferencer struct {
	ImageURL string
}

func (ir *ImageReferencer) ToSQLC() (*pgtype.Text, error) {
	var imageURL pgtype.Text
	if err := imageURL.Scan(ir.ImageURL); err != nil {
		return nil, err
	}
	return &imageURL, nil
}
