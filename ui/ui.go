// Package ui is the package for the main UI of the simon-says game.
package ui

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// all possible Events to communicate between components.
const (
	EventClick       = "click"
	EventSimonSays   = "playSequence"
	EventPlayButton  = "play%d"
	EventNewGame     = "newGame"
	EventStateChange = "stateChange"
)

// NewUI is the factory for the main UI component.
func NewUI() *UI {
	return &UI{}
}

// UI is the main UI component of the game.
type UI struct {
	app.Compo

	Text            string
	updateAvailable bool

	showNewGame bool
	playGame    bool
	showStats   bool
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (u *UI) OnMount(ctx app.Context) {
	ctx.Handle(EventStateChange, u.handleStateChange)
}

// OnAppUpdate satisfies the app.AppUpdater interface. It is called when the app
// is updated in background.
func (u *UI) OnAppUpdate(ctx app.Context) {
	u.updateAvailable = ctx.AppUpdateAvailable()
}

// Render implements the interface for go-app to render the component.
func (u *UI) Render() app.UI {
	if u.Text == "" {
		u.Text = "Start a New Game"
	}
	menu := &menu{}
	gameStateText := app.Div().Class("game-state").Text(u.Text)
	gameField := app.Div().Class("game-field")

	firstButton := newButton(0)
	secondButton := newButton(1)
	thirdButton := newButton(2)
	fourthButton := newButton(3)

	gameField.Body(
		firstButton,
		secondButton,
		thirdButton,
		fourthButton,
	)

	return app.Div().Body(
		app.A().Href("/stats").Body(app.Img().Src("web/stats.png").Style("height", "29px").Style("width", "29px")),
		menu,
		gameStateText,
		gameField,
		app.If(u.updateAvailable,
			app.Button().Class("simon-button", "update").
				Body(app.Span().Text("Update!")).
				OnClick(u.onUpdateClick),
		),
	)
}

func (u *UI) onUpdateClick(ctx app.Context, _ app.Event) {
	// Reloads the page to display the modifications.
	ctx.Reload()
}

func (u *UI) handleStateChange(_ app.Context, a app.Action) {
	txt, ok := a.Value.(string)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	u.Text = txt
}
