package release

type ReleaseType int

const (
	Album ReleaseType = iota
	EP
	Single
	Compilation
)

var releaseTypesStrings []string = []string{"album", "ep", "single", "compilation"}

func (rt *ReleaseType) ToString() string {
	return releaseTypesStrings[*rt]
}
