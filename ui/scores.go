package ui

import (
	"fmt"
	"math"
	"reflect"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var LocalStorageScores = "scores"

type Scores struct {
	Basic   map[Difficulty]Score
	Endless map[int]int
}

type Score struct {
	Win  int
	Loss int
}

type scoreBoard struct {
	app.Compo

	easyWinRatio     float64
	mediumWinRatio   float64
	hardWinRatio     float64
	endlessHighscore int
}

func (g *scoreBoard) OnMount(ctx app.Context) {
	ctx.Handle(EventScoreUpdate, g.HandleScoreUpdate)
	s := &Scores{}
	ctx.LocalStorage().Get(LocalStorageScores, s)
	if reflect.DeepEqual(s, &Scores{}) {
		s = &Scores{Basic: map[Difficulty]Score{
			Easy:   {},
			Medium: {},
			Hard:   {},
		}, Endless: map[int]int{}}
	}

	g.storeScores(s)
}

func (g *scoreBoard) storeScores(s *Scores) {
	g.easyWinRatio = float64(s.Basic[Easy].Win) / float64(s.Basic[Easy].Win+s.Basic[Easy].Loss)
	g.mediumWinRatio = float64(s.Basic[Medium].Win) / float64(s.Basic[Medium].Win+s.Basic[Medium].Loss)
	g.hardWinRatio = float64(s.Basic[Hard].Win) / float64(s.Basic[Hard].Win+s.Basic[Hard].Loss)

	max := 0
	for score := range s.Endless {
		if score > max {
			max = score
		}
	}
	g.endlessHighscore = max
}

func (s *scoreBoard) Render() app.UI {
	easyText := "no game yet"
	if !math.IsNaN(s.easyWinRatio) {
		easyText = fmt.Sprintf("%.1f%%", s.easyWinRatio*100)
	}
	mediumText := "no game yet"
	if !math.IsNaN(s.mediumWinRatio) {
		mediumText = fmt.Sprintf("%.1f%%", s.mediumWinRatio*100)
	}
	hardText := "no game yet"
	if !math.IsNaN(s.hardWinRatio) {
		hardText = fmt.Sprintf("%.1f%%", s.hardWinRatio*100)
	}

	return app.Table().Class("scores").Body(
		app.Tr().Body(app.Td().Text("Easy"), app.Td().Text(easyText)),
		app.Tr().Body(app.Td().Text("Medium"), app.Td().Text(mediumText)),
		app.Tr().Body(app.Td().Text("Hard"), app.Td().Text(hardText)),
		app.Tr().Body(app.Td().Text("Endless"), app.Td().Text(s.endlessHighscore)),
	)
}

func (s *scoreBoard) HandleScoreUpdate(ctx app.Context, a app.Action) {
	scores, ok := a.Value.(*Scores)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	s.storeScores(scores)
}
