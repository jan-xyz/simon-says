package ui

import (
	"fmt"
	"math"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Stats struct {
	app.Compo

	scores storage.Scores
}

// OnMount implements the Mounter interface to run this on mounting the component.
func (s *Stats) OnMount(ctx app.Context) {
	ctx.Handle(storage.EventScoreUpdate, s.HandleScoreUpdate)
	scores := storage.LoadScores(ctx)
	s.scores = scores
}

// Render implements the interface for go-app to render the component.
func (s *Stats) Render() app.UI {
	scores := s.scores
	easyWinRatio := float64(scores.Basic[storage.Easy].Win) / float64(scores.Basic[storage.Easy].Win+scores.Basic[storage.Easy].Loss)
	mediumWinRatio := float64(scores.Basic[storage.Medium].Win) / float64(scores.Basic[storage.Medium].Win+scores.Basic[storage.Medium].Loss)
	hardWinRatio := float64(scores.Basic[storage.Hard].Win) / float64(scores.Basic[storage.Hard].Win+scores.Basic[storage.Hard].Loss)

	max := 0
	for score := range scores.Endless {
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

	chart := &barChart{Bar: newBarChart(s.scores.Endless)}
	stats := app.Table().Class("scores").Body(
		app.Tr().Body(app.Td().Text("Easy"), app.Td().Text(easyText)),
		app.Tr().Body(app.Td().Text("Medium"), app.Td().Text(mediumText)),
		app.Tr().Body(app.Td().Text("Hard"), app.Td().Text(hardText)),
		app.Tr().Body(app.Td().Text("Endless"), app.Td().Text(max)),
	)
	return app.Span().Body(
		app.A().Href("/").Body(app.Img().Src("web/stats.png").Style("height", "29px").Style("width", "29px")),
		stats,
		chart,
	)
}

func (s *Stats) HandleScoreUpdate(_ app.Context, a app.Action) {
	scores, ok := a.Value.(storage.Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	s.scores = scores
}
