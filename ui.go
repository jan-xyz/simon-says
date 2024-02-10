package main

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	eventClick       events = "click"
	eventSimonSays   events = "playSequence"
	eventPlayButton  events = "play%d"
	eventNewGame     events = "newGame"
	eventStateChange events = "stateChange"
)

func NewUI() *ui {
	return &ui{}
}

type ui struct {
	app.Compo

	Text               string
	selectedDifficulty difficulty
}

func (g *ui) OnMount(ctx app.Context) {
	ctx.Handle(eventStateChange, g.handleStateChange)
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
		gameStateText,
	)

	newGameButton := app.Button().
		Class("simon-button", "new-game").
		Body(app.Span().Text("New Game")).
		OnClick(func(ctx app.Context, _ app.Event) {
			ctx.NewActionWithValue(eventNewGame, g.selectedDifficulty)
		})
	difficultySetting := app.Div().ID("difficulty").Body(
		app.Label().Body(
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", easy).Value(easy).Checked(true).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				d, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("failed parsing", val)
					return
				}
				g.selectedDifficulty = difficulty(d)
			}),
			app.Text("easy"),
		),
		app.Label().Body(
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", medium).Value(medium).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				d, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("failed parsing", val)
					return
				}
				g.selectedDifficulty = difficulty(d)
			}),
			app.Text("medium"),
		),
		app.Label().Body(
			app.Input().Type("radio").Name("difficulty-setting").ID("difficulty%d", hard).Value(hard).OnClick(func(ctx app.Context, _ app.Event) {
				val := ctx.JSSrc().Get("value").String()
				d, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("failed parsing", val)
					return
				}
				g.selectedDifficulty = difficulty(d)
			}),
			app.Text("hard"),
		),
	)

	return app.Div().Class("fill").Body(
		newGameButton,
		difficultySetting,
		gameField,
	)
}

func (b *ui) handleStateChange(ctx app.Context, a app.Action) {
	state, ok := a.Value.(gameState)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	txt := ""
	switch state {
	case gameStateNoGame:
		txt = "Start a New Game"
	case gameStatePlayerSays:
		txt = "Your turn"
	case gameStateSimonSays:
		txt = "Simon says..."
	case gameStateLost:
		txt = "You Lost. Start a New Game"
	case gameStateWon:
		txt = "You Won. Start a New Game"
	}
	ctx.Dispatch(func(_ app.Context) {
		b.Text = txt
	})
}
