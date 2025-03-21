package command

import "github.com/google/uuid"

type RetrieveCommand struct {
	Id uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" validate:"required,uuid`
}
