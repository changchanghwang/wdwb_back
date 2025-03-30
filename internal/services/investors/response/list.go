package response

type InvestorListResponse struct {
	Items []*InvestorRetrieveResponse `json:"items"`
	Count int                         `json:"count"`
}
