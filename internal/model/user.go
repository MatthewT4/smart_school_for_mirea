package model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID         uuid.UUID `db:"id" json:"uuid"`
	Email        string
	HashPassword string `db:"password"`
}
