package storage

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const localStorageDifficulty = "difficulty"

// Difficulty represents the games difficulty for the statistics.
type Difficulty string

// List of possible difficulties.
const (
	Easy    Difficulty = "easy"
	Medium  Difficulty = "medium"
	Hard    Difficulty = "hard"
	Endless Difficulty = "endless"
)

// LoadDifficulty loads the difficulty from local storage.
// If the value is not supported, it defaults to [Easy]
func LoadDifficulty(ctx app.Context) Difficulty {
	var val string
	ctx.LocalStorage().Get(localStorageDifficulty, &val)
	d := Easy
	switch val {
	case string(Easy):
		d = Easy
	case string(Medium):
		d = Medium
	case string(Hard):
		d = Hard
	case string(Endless):
		d = Endless
	}
	return d
}

// SetDifficulty stores the provided value. If the value is not supported,
// it defaults to [Easy]
func SetDifficulty(ctx app.Context, val string) Difficulty {
	d := Easy
	switch val {
	case string(Easy):
		d = Easy
	case string(Medium):
		d = Medium
	case string(Hard):
		d = Hard
	case string(Endless):
		d = Endless
	}
	ctx.LocalStorage().Set(localStorageDifficulty, d)
	return d
}
