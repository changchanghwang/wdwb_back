package commands

type RankCommand struct {
	Year    int `json:"year" validate:"required"`
	Quarter int `json:"quarter"`
}
