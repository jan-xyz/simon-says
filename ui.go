package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	click events = "click"
)

func NewSimonSaysUI() *simonSaysUI {
	return &simonSaysUI{}
}

type simonSaysUI struct {
	app.Compo
}

func (h *simonSaysUI) Render() app.UI {
	t := app.Div().Class("game-field")

	t.Body(
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.onClickFirst),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.onClickSecond),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.onClickThird),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.onClickFourth),
	)

	return app.Div().Class("fill", "background").Body(
		t,
	)
}

func (h *simonSaysUI) onClickFirst(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, int64(0))
}

func (h *simonSaysUI) onClickSecond(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, int64(1))
}

func (h *simonSaysUI) onClickThird(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, int64(2))
}

func (h *simonSaysUI) onClickFourth(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, int64(3))
}
