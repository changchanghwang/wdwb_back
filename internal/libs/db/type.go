package db

type FindOptions struct {
	Offset  int
	Limit   int
	GroupBy string
}

type OrderOptions struct {
	OrderBy   string
	Direction string
}
