package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type button struct {
	app.Compo

	id int64
}

func NewButton(id int64) *button {
	return &button{
		id: id,
	}
}

func (b *button) Render() app.UI {
	return app.Button().
		Class("simon-button").
		Body(app.Span().Text("")).
		ID("button%d", b.id).
		OnClick(func(ctx app.Context, _ app.Event) {
			ctx.NewActionWithValue(click, b.id)
		})
}
