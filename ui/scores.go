package ui

import (
	"fmt"
	"math"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type scoreBoard struct {
	app.Compo

	easyWinRatio     float64
	mediumWinRatio   float64
	hardWinRatio     float64
	endlessHighscore int
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (b *scoreBoard) OnMount(ctx app.Context) {
	ctx.Handle(storage.EventScoreUpdate, b.HandleScoreUpdate)
	s := storage.LoadScores(ctx)
	b.storeScores(s)
}

func (b *scoreBoard) storeScores(s storage.Scores) {
	b.easyWinRatio = float64(s.Basic[storage.Easy].Win) / float64(s.Basic[storage.Easy].Win+s.Basic[storage.Easy].Loss)
	b.mediumWinRatio = float64(s.Basic[storage.Medium].Win) / float64(s.Basic[storage.Medium].Win+s.Basic[storage.Medium].Loss)
	b.hardWinRatio = float64(s.Basic[storage.Hard].Win) / float64(s.Basic[storage.Hard].Win+s.Basic[storage.Hard].Loss)

	max := 0
	for score := range s.Endless {
		if score > max {
			max = score
		}
	}
	b.endlessHighscore = max
}

// Render implements the interface for go-app to render the component.
func (b *scoreBoard) Render() app.UI {
	easyText := "no game yet"
	if !math.IsNaN(b.easyWinRatio) {
		easyText = fmt.Sprintf("%.1f%%", b.easyWinRatio*100)
	}
	mediumText := "no game yet"
	if !math.IsNaN(b.mediumWinRatio) {
		mediumText = fmt.Sprintf("%.1f%%", b.mediumWinRatio*100)
	}
	hardText := "no game yet"
	if !math.IsNaN(b.hardWinRatio) {
		hardText = fmt.Sprintf("%.1f%%", b.hardWinRatio*100)
	}

	chart := &GoAppBar{Class: "chart1-cls"}
	return app.Table().Class("scores").Body(
		app.Tr().Body(app.Td().Text("Easy"), app.Td().Text(easyText)),
		app.Tr().Body(app.Td().Text("Medium"), app.Td().Text(mediumText)),
		app.Tr().Body(app.Td().Text("Hard"), app.Td().Text(hardText)),
		app.Tr().Body(app.Td().Text("Endless"), app.Td().Text(b.endlessHighscore)),
		app.Button().Text("update").
			OnClick(func(ctx app.Context, _ app.Event) {
				chart.UpdateConfig(ctx, newChart())
			}),
		chart,
	)
}

func (b *scoreBoard) HandleScoreUpdate(_ app.Context, a app.Action) {
	scores, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	b.storeScores(scores)
}
