package main

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type menu struct {
	app.Compo
	selectedDifficulty difficulty
}

func NewMenu() *menu {
	return &menu{}
}

func (g *menu) Render() app.UI {
	return app.Div().Body(

		app.Button().
			Class("simon-button", "new-game").
			Body(app.Span().Text("New Game")).
			OnClick(func(ctx app.Context, _ app.Event) {
				ctx.NewActionWithValue(eventNewGame, g.selectedDifficulty)
			}),
		app.Div().ID("difficulty").Body(
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
		),
	)
}