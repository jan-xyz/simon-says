package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func NewSimonSays() *simonSays {
	s := &simonSays{}
	return s
}

type simonSays struct {
	app.Compo
}

func (h *simonSays) Render() app.UI {
	t := app.Div().Class("game-field")

	t.Body(
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.firstButtonClicked),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.secondButtonClicked),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.thirdButtonClicked),
		app.Button().
			Class("simon-button").
			Body(app.Span().Text("")).OnClick(h.fourthButtonClicked),
	)

	return app.Div().Class("fill", "background").Body(
		t,
	)
}

func (h *simonSays) firstButtonClicked(ctx app.Context, e app.Event) {
	fmt.Println("clicked first button")
}

func (h *simonSays) secondButtonClicked(ctx app.Context, e app.Event) {
	fmt.Println("clicked second button")
}

func (h *simonSays) thirdButtonClicked(ctx app.Context, e app.Event) {
	fmt.Println("clicked third button")
}

func (h *simonSays) fourthButtonClicked(ctx app.Context, e app.Event) {
	fmt.Println("clicked fourth button")
}
