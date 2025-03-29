package response

import "github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"

type ListResponse struct {
	//TODO: specific fields
	Items []*domain.Holding `json:"items"`
	Count int               `json:"count"`
}
