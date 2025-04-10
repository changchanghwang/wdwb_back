package commands

type RankCommand struct {
	Year    int `json:"year" validate:"required"` // 2024
	Quarter int `json:"quarter"`                  // 1, 2, 3, 4 optional
}
