package ui

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	EventClick       events = "click"
	EventSimonSays   events = "playSequence"
	EventPlayButton  events = "play%d"
	EventNewGame     events = "newGame"
	EventStateChange events = "stateChange"
)

func NewUI() *ui {
	return &ui{}
}

type ui struct {
	app.Compo

	Text            string
	updateAvailable bool
}

func (g *ui) OnMount(ctx app.Context) {
	ctx.Handle(EventStateChange, g.handleStateChange)
}

// OnAppUpdate satisfies the app.AppUpdater interface. It is called when the app
// is updated in background.
func (g *ui) OnAppUpdate(ctx app.Context) {
	g.updateAvailable = ctx.AppUpdateAvailable()
}

func (g *ui) Render() app.UI {
	if g.Text == "" {
		g.Text = "Start a New Game"
	}
	gameField := app.Div().Class("game-field")

	firstButton := NewButton(0)
	secondButton := NewButton(1)
	thirdButton := NewButton(2)
	fourthButton := NewButton(3)

	gameStateText := app.Div().Class("game-state").Text(g.Text)

	gameField.Body(
		firstButton,
		secondButton,
		thirdButton,
		fourthButton,
	)

	menu := NewMenu()

	return app.Div().Class("fill").Body(
		menu,
		gameStateText,
		gameField,
		app.If(g.updateAvailable,
			app.Button().
				Text("Update!").
				OnClick(g.onUpdateClick),
		),
	)
}

func (g *ui) onUpdateClick(ctx app.Context, e app.Event) {
	// Reloads the page to display the modifications.
	ctx.Reload()
}

func (b *ui) handleStateChange(ctx app.Context, a app.Action) {
	txt, ok := a.Value.(string)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	ctx.Dispatch(func(_ app.Context) {
		b.Text = txt
	})
}
