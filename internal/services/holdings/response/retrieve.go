package response

import "github.com/google/uuid"

type RetrieveResponse struct {
	Id         uuid.UUID `json:"id"`
	InvestorId uuid.UUID `json:"investorId"`
	Name       string    `json:"name"`
	Year       int       `json:"year"`
	Quarter    int       `json:"quarter"`
	Value      int       `json:"value"`
	Shares     int       `json:"shares"`
}
