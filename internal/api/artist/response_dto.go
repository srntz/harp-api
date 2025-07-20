package artist

import "github.com/srntz/harp-api/internal/domain/artist"

type ResponseDTO struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	ImageURL string `json:"image_url"`
}

func ResponseDTOFromDomainModel(a *artist.Artist) *ResponseDTO {
	return &ResponseDTO{
		Name:     a.Name,
		Bio:      a.Bio,
		ImageURL: a.ImageURL,
	}
}
