package response

type ListResponse struct {
	Items []*RetrieveResponse `json:"items"`
	Count int                 `json:"count"`
}
