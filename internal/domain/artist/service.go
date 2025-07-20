package artist

type ArtistService struct {
	r IArtistRepository
}

func NewArtistService(repository IArtistRepository) *ArtistService {
	return &ArtistService{
		r: repository,
	}
}

func (s *ArtistService) GetByID(id string) (*Artist, error) {
	artist, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
