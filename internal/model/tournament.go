package model

import "github.com/google/uuid"

type Tournament struct {
	ID     uuid.UUID
	Name   string
	Status string
	Type   string
}
