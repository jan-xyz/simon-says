package ui

import (
	"fmt"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type menu struct {
	app.Compo
	selectedDifficulty storage.Difficulty
	selectedSpeed      storage.Speed
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (g *menu) OnMount(ctx app.Context) {
	d := storage.LoadDifficulty(ctx)
	g.selectedDifficulty = d
	s := storage.LoadSpeed(ctx)
	g.selectedSpeed = s
}

// Render implements the interface for go-app to render the component.
func (g *menu) Render() app.UI {
	modes := []app.UI{}
	speedPanel := []app.UI{}
	for _, mode := range []storage.Difficulty{storage.Easy, storage.Medium, storage.Hard, storage.Endless} {
		id := fmt.Sprintf("difficulty%s", mode)
		input := app.Input().Type("radio").Name("difficulty-setting").ID(id).Value(mode).OnClick(g.storeDifficulty)
		label := app.Label().For("difficulty%s", mode).Body(app.Span().Text(mode))
		if g.selectedDifficulty == mode {
			input.Checked(true)
		}
		modes = append(modes, input, label)
	}
	for _, speed := range []storage.Speed{storage.Slow, storage.Normal, storage.Fast} {
		id := fmt.Sprintf("speed%s", speed)
		input := app.Input().Type("radio").Name("speed-setting").ID(id).Value(speed).OnClick(g.storeSpeed)
		label := app.Label().For("speed%s", speed).Body(app.Span().Text(speed))
		if g.selectedSpeed == speed {
			input.Checked(true)
		}
		speedPanel = append(speedPanel, input, label)
	}
	t := app.Div().Body(
		app.Input().Type("image").Src("web/stats.png").Style("height", "29px").Style("width", "29px"),
		app.Button().
			Class("simon-button", "new-game").
			Body(app.Span().Text("New Game")).
			OnClick(g.startNewGame),
		app.Div().Class("difficulty-settings").Body(
			modes...,
		),
		app.Div().Class("speed-settings").Body(
			speedPanel...,
		),
	)
	return t
}

func (g *menu) storeDifficulty(ctx app.Context, _ app.Event) {
	val := ctx.JSSrc().Get("value").String()
	d := storage.SetDifficulty(ctx, val)
	g.selectedDifficulty = d
}

func (g *menu) storeSpeed(ctx app.Context, _ app.Event) {
	val := ctx.JSSrc().Get("value").String()
	s := storage.SetSpeed(ctx, val)
	g.selectedSpeed = s
}

// NewGameSettings are the settings of a new game that gets published to the logic
type NewGameSettings struct {
	Speed      storage.Speed
	Difficulty storage.Difficulty
}

func (g *menu) startNewGame(ctx app.Context, _ app.Event) {
	ctx.NewActionWithValue(EventNewGame, NewGameSettings{Speed: g.selectedSpeed, Difficulty: g.selectedDifficulty})
}
