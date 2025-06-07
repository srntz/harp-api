package user

import "github.com/srntz/harp-api/internal/domain/shared"

type User struct {
	shared.Entity
	username string
	bio      string
}
