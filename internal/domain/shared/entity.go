package shared

type EntityIndentifier int

type Entity struct {
	id         EntityIndentifier
	created_at int
	updated_at int
}
