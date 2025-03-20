package response

type RetrieveResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CompanyName  string `json:"companyName"`
	Cik          string `json:"cik"`
	HoldingValue int    `json:"holdingValue"`
}
