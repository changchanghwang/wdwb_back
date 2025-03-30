package response

import "github.com/google/uuid"

type InvestorRetrieveResponse struct {
	Id           uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"` // investor id
	Name         string    `json:"name" example:"John Doe"`                           // investor name
	CompanyName  string    `json:"companyName" example:"Company Name"`                // investor company name
	Cik          string    `json:"cik" example:"1234567890"`                          // investor company index key
	HoldingValue int       `json:"holdingValue" example:"1000000"`                    // total value of holdings
}
