package main

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

func NewButton(id int64) *button {
	return &button{
		id: id,
	}
}

func (b *button) OnMount(ctx app.Context) {
	ctx.Handle(fmt.Sprintf(playButton, b.id), b.handleActivate)
}

func (b *button) Render() app.UI {
	e := app.Button().
		Class("simon-button").
		Body(app.Span().Text("")).
		ID("button%d", b.id).
		OnClick(func(ctx app.Context, _ app.Event) {
			ctx.NewActionWithValue(click, b.id)
		})
	if b.Active {
		e.Class("active")
	}
	return e
}

func (b *button) handleActivate(ctx app.Context, a app.Action) {
	fmt.Println("playing", b.id)
	ctx.Dispatch(func(_ app.Context) {
		b.Active = true
	})
	ctx.After(800*time.Millisecond, func(_ app.Context) {
		ctx.Dispatch(func(_ app.Context) {
			b.Active = false
		})
	})
}
