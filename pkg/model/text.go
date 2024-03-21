package model

import "github.com/google/uuid"

type Translation struct {
	ID      uuid.UUID
	Lang    string
	Content string
}
