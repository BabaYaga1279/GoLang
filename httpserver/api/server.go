package api

import (
	"github.com/google/uuid"
)

type Item struct {
	ID uuid.UUID `json:"ID"`
}
