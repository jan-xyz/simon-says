package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	click events = "click"
)

func NewSimonSays() *simonSays {
	return &simonSays{}
}

type simonSays struct {
	app.Compo
}

func (h *simonSays) Render() app.UI {
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

func (h *simonSays) onClickFirst(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, 0)
}

func (h *simonSays) onClickSecond(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, 1)
}

func (h *simonSays) onClickThird(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, 2)
}

func (h *simonSays) onClickFourth(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue(click, 3)
}
