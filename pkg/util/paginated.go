package util

type PaginatedResponse[T comparable] struct {
	Items []T `json:"items"`
	Count int `json:"count"`
}

func Paginated[T comparable](items []T, count int) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Items: items,
		Count: count,
	}
}
