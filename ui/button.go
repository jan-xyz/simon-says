package ui

import (
	"fmt"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type button struct {
	app.Compo

	id     int64
	Active bool
}

func newButton(id int64) *button {
	return &button{
		id: id,
	}
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (b *button) OnMount(ctx app.Context) {
	ctx.Handle(fmt.Sprintf(EventPlayButton, b.id), b.handlePlayButton)
}

// Render implements the interface for go-app to render the component.
func (b *button) Render() app.UI {
	id := fmt.Sprintf("button%d", b.id)
	e := app.Button().
		Class("simon-button", "game-button").
		Body(app.Span().Text("")).
		ID(id).
		OnClick(b.handleClick)
	if b.Active {
		e.Class("active")
	}
	return e
}

func (b *button) handleClick(ctx app.Context, _ app.Event) {
	ctx.NewActionWithValue(EventClick, b.id)
	// Needs a short delay because it doesn't work if it is done directly on click
	// This is done to also have a click animation on touch devices.
	ctx.After(50*time.Millisecond, func(_ app.Context) {
		b.Active = true
		ctx.After(300*time.Millisecond, func(_ app.Context) {
			b.Active = false
		})
	})
}

func (b *button) handlePlayButton(ctx app.Context, _ app.Action) {
	b.Active = true
	ctx.After(800*time.Millisecond, func(_ app.Context) {
		ctx.Dispatch(func(_ app.Context) {
			b.Active = false
		})
	})
}
