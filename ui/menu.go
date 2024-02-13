package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type (
	Difficulty string
)

const (
	Easy    Difficulty = "easy"
	Medium  Difficulty = "medium"
	Hard    Difficulty = "hard"
	Endless Difficulty = "endless"
)

const (
	localStorageDifficulty = "difficulty"
)

type menu struct {
	app.Compo
	selectedDifficulty Difficulty
}

func NewMenu() *menu {
	return &menu{}
}

func (g *menu) OnMount(ctx app.Context) {
	var d Difficulty
	ctx.LocalStorage().Get(localStorageDifficulty, &d)
	if d == "" {
		d = Easy
	}
	g.selectedDifficulty = d
}

func (g *menu) Render() app.UI {
	modes := []app.UI{}
	for _, mode := range []Difficulty{Easy, Medium, Hard, Endless} {
		input := app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", mode).Value(mode).OnClick(g.storeValue)
		label := app.Label().For("difficulty%d", mode).Body(app.Span().Text(mode))
		if g.selectedDifficulty == mode {
			input.Checked(true)
		}
		modes = append(modes, input, label)
	}
	t := app.Div().Body(

		app.Button().
			Class("simon-button", "new-game").
			Body(app.Span().Text("New Game")).
			OnClick(g.startNewGame),
		app.Div().Class("difficulty").Body(
			modes...,
		),
	)
	return t
}

func (g *menu) storeValue(ctx app.Context, _ app.Event) {
	val := ctx.JSSrc().Get("value").String()
	ctx.LocalStorage().Set(localStorageDifficulty, val)
	g.selectedDifficulty = Difficulty(val)
}

func (g *menu) startNewGame(ctx app.Context, _ app.Event) {
	var d Difficulty
	ctx.LocalStorage().Get(localStorageDifficulty, &d)
	if d == "" {
		d = Easy
	}
	ctx.NewActionWithValue(EventNewGame, d)
}
