package response

import "github.com/google/uuid"

type HoldingRetrieveResponse struct {
	Id         string    `json:"id" example:"0001067983"`                                   // holding cik
	InvestorId uuid.UUID `json:"investorId" example:"123e4567-e89b-12d3-a456-426614174000"` // investor id
	Name       string    `json:"name" example:"Company Name"`                               // holding name
	Year       int       `json:"year" example:"2021"`                                       // holding year
	Quarter    int       `json:"quarter" example:"1"`                                       // holding quarter
	Value      int       `json:"value" example:"1000000"`                                   // total amount of holding value
	Shares     int       `json:"shares" example:"1000"`                                     // number of stock shares
	Translated bool      `json:"translated" example:"true"`                                 // whether the holding name is translated
}
