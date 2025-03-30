package response

type HoldingListResponse struct {
	Items []*HoldingRetrieveResponse `json:"items"`
	Count int                        `json:"count"`
}
