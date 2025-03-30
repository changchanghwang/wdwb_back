package commands

import "github.com/google/uuid"

type ListCommand struct {
	InvestorId uuid.UUID `json:"investorId" validate:"required,uuid"`
	Year       int       `json:"year" validate:"required,number"`
	Quarter    int       `json:"quarter" validate:"required,number"`
	Page       int       `json:"page" validate:"number"`
	Limit      int       `json:"limit" validate:"number"`
}
