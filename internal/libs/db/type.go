package db

type FindOptions struct {
	Offset int
	Limit  int
}

type OrderOptions struct {
	OrderBy   string
	Direction string
}
