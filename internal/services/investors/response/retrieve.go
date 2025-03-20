package response

type RetrieveResponse struct {
	Id           string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name         string `json:"name" example:"John Doe"`
	CompanyName  string `json:"companyName" example:"Company Name"`
	Cik          string `json:"cik" example:"1234567890"`
	HoldingValue int    `json:"holdingValue" example:"1000000"`
}
