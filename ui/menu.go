package ui

import (
	"fmt"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type menu struct {
	app.Compo
	selectedDifficulty storage.Difficulty
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (g *menu) OnMount(ctx app.Context) {
	d := storage.LoadDifficulty(ctx)
	g.selectedDifficulty = d
}

// Render implements the interface for go-app to render the component.
func (g *menu) Render() app.UI {
	modes := []app.UI{}
	for _, mode := range []storage.Difficulty{storage.Easy, storage.Medium, storage.Hard, storage.Endless} {
		id := fmt.Sprintf("difficulty%s", mode)
		input := app.Input().Type("radio").Name("difficulty-setting").ID(id).Value(mode).OnClick(g.storeValue)
		label := app.Label().For("difficulty%s", mode).Body(app.Span().Text(mode))
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
	d := storage.SetDifficulty(ctx, val)
	g.selectedDifficulty = d
}

func (g *menu) startNewGame(ctx app.Context, _ app.Event) {
	ctx.NewActionWithValue(EventNewGame, g.selectedDifficulty)
}
