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

type menu struct {
	app.Compo
	selectedDifficulty Difficulty
}

func NewMenu() *menu {
	return &menu{selectedDifficulty: Easy}
}

func (g *menu) Render() app.UI {
	return app.Div().Body(

		app.Button().
			Class("simon-button", "new-game").
			Body(app.Span().Text("New Game")).
			OnClick(func(ctx app.Context, _ app.Event) {
				ctx.NewActionWithValue(EventNewGame, g.selectedDifficulty)
			}),
		app.Div().Class("difficulty").Body(
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", Easy).Value(Easy).Checked(true).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				g.selectedDifficulty = Difficulty(val)
			}),
			app.Label().For("difficulty%d", Easy).Body(app.Span().Text("easy")),
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", Medium).Value(Medium).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				g.selectedDifficulty = Difficulty(val)
			}),
			app.Label().For("difficulty%d", Medium).Body(app.Span().Text("medium")),
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", Hard).Value(Hard).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				g.selectedDifficulty = Difficulty(val)
			}),
			app.Label().For("difficulty%d", Hard).Body(app.Span().Text("hard")),
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", Endless).Value(Endless).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				g.selectedDifficulty = Difficulty(val)
			}),
			app.Label().For("difficulty%d", Endless).Body(app.Span().Text("endless")),
		),
	)
}
