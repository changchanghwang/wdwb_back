package commands

import "github.com/google/uuid"

type ListCommand struct {
	InvestorId uuid.UUID `query:"investorId"`
	Page       int       `query:"page"`
	Limit      int       `query:"limit"`
}
