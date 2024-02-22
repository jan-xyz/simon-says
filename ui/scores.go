package ui

import (
	"fmt"
	"math"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type scoreBoard struct {
	app.Compo

	scores storage.Scores
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (b *scoreBoard) OnMount(ctx app.Context) {
	ctx.Handle(storage.EventScoreUpdate, b.HandleScoreUpdate)
	s := storage.LoadScores(ctx)
	b.scores = s
}

// Render implements the interface for go-app to render the component.
func (b *scoreBoard) Render() app.UI {
	s := b.scores
	easyWinRatio := float64(s.Basic[storage.Easy].Win) / float64(s.Basic[storage.Easy].Win+s.Basic[storage.Easy].Loss)
	mediumWinRatio := float64(s.Basic[storage.Medium].Win) / float64(s.Basic[storage.Medium].Win+s.Basic[storage.Medium].Loss)
	hardWinRatio := float64(s.Basic[storage.Hard].Win) / float64(s.Basic[storage.Hard].Win+s.Basic[storage.Hard].Loss)

	max := 0
	for score := range s.Endless {
		if score > max {
			max = score
		}
	}
	easyText := "no game yet"
	if !math.IsNaN(easyWinRatio) {
		easyText = fmt.Sprintf("%.1f%%", easyWinRatio*100)
	}
	mediumText := "no game yet"
	if !math.IsNaN(mediumWinRatio) {
		mediumText = fmt.Sprintf("%.1f%%", mediumWinRatio*100)
	}
	hardText := "no game yet"
	if !math.IsNaN(hardWinRatio) {
		hardText = fmt.Sprintf("%.1f%%", hardWinRatio*100)
	}

	chart := &barChart{Bar: newBarChart(b.scores.Endless)}
	scores := app.Table().Class("scores").Body(
		app.Tr().Body(app.Td().Text("Easy"), app.Td().Text(easyText)),
		app.Tr().Body(app.Td().Text("Medium"), app.Td().Text(mediumText)),
		app.Tr().Body(app.Td().Text("Hard"), app.Td().Text(hardText)),
		app.Tr().Body(app.Td().Text("Endless"), app.Td().Text(max)),
	)
	return app.Span().Body(scores, chart)
}

func (b *scoreBoard) HandleScoreUpdate(_ app.Context, a app.Action) {
	s, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	b.scores = s
}
