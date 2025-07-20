package artist

type IArtistRepository interface {
	GetByID(id string) (*Artist, error)
}
