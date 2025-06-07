package review

import "github.com/srntz/harp-api/internal/domain/shared"

type Review struct {
	shared.Entity
	header  string
	body    string
	user_id shared.EntityIndentifier
}
