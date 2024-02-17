package storage

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// EventScoreUpdate is fired when a score is updated.
const EventScoreUpdate = "scoreUpdate"

const localStorageScores = "scores"

type Scores struct {
	Basic   map[Difficulty]Score
	Endless map[int]int
}

type Difficulty string

const (
	Easy    Difficulty = "easy"
	Medium  Difficulty = "medium"
	Hard    Difficulty = "hard"
	Endless Difficulty = "endless"
)

type Score struct {
	Win  int
	Loss int
}

func newScores() Scores {
	return Scores{Basic: map[Difficulty]Score{
		Easy:   {},
		Medium: {},
		Hard:   {},
	}, Endless: map[int]int{}}
}

func IncrementEasyLoss(ctx app.Context) {
	IncrementLoss(ctx, Easy)
}

func IncrementMediumLoss(ctx app.Context) {
	IncrementLoss(ctx, Medium)
}

func IncrementHardLoss(ctx app.Context) {
	IncrementLoss(ctx, Hard)
}

func IncrementLoss(ctx app.Context, d Difficulty) {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	if d != Endless {
		f := s.Basic[d]
		f.Loss++
		s.Basic[d] = f
	}
	ctx.LocalStorage().Set(localStorageScores, s)
	ctx.NewActionWithValue(EventScoreUpdate, s)
}

func UpdateEndless(ctx app.Context, stage int) {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	s.Endless[stage]++
	ctx.LocalStorage().Set(localStorageScores, s)
	ctx.NewActionWithValue(EventScoreUpdate, s)
}

func IncrementEasyWin(ctx app.Context) {
	IncrementWin(ctx, Easy)
}

func IncrementMediumWin(ctx app.Context) {
	IncrementWin(ctx, Medium)
}

func IncrementHardWin(ctx app.Context) {
	IncrementWin(ctx, Hard)
}

func IncrementWin(ctx app.Context, d Difficulty) {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	f := s.Basic[d]
	f.Win++
	s.Basic[d] = f
	ctx.LocalStorage().Set(localStorageScores, s)
	ctx.NewActionWithValue(EventScoreUpdate, s)
}

func LoadScores(ctx app.Context) Scores {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	return s
}
