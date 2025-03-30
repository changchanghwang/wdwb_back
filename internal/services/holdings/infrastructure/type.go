package infrastructure

import "github.com/google/uuid"

type HoldingQueryConditions struct {
	InvestorIds []uuid.UUID
	Years       []int
	Quarters    []int
}
