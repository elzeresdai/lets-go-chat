package models

import "github.com/google/uuid"

type UserDB struct {
	ID   uuid.UUID `db:"ID"`
	Name string    `db:"name"`
	Hash string    `db:"hash"`
}
