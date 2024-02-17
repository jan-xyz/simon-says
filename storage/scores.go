// Package storage is a library to make interfacing with the web storage simpler
// and provide a unified way to deal with it.
package storage

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// EventScoreUpdate is fired when a score is updated.
const EventScoreUpdate = "scoreUpdate"

const localStorageScores = "scores"

// Scores is the representation of scores for ranking and statistics
type Scores struct {
	Basic   map[Difficulty]Score
	Endless map[int]int
}

// Difficulty represents the games difficulty for the statistics.
type Difficulty string

// List of possible difficulties.
const (
	Easy    Difficulty = "easy"
	Medium  Difficulty = "medium"
	Hard    Difficulty = "hard"
	Endless Difficulty = "endless"
)

// Score is the individual scores for the basic difficulties. Endless doesn't have Wins that
// is why it isn't tracked with Win/Loss Scores.
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

// IncrementEasyLoss increments the losses for an easy game.
func IncrementEasyLoss(ctx app.Context) {
	incrementLoss(ctx, Easy)
}

// IncrementMediumLoss increments the losses for a medium game.
func IncrementMediumLoss(ctx app.Context) {
	incrementLoss(ctx, Medium)
}

// IncrementHardLoss increments the losses for a hard game.
func IncrementHardLoss(ctx app.Context) {
	incrementLoss(ctx, Hard)
}

func incrementLoss(ctx app.Context, d Difficulty) {
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

// UpdateEndless tracks the score for an endless game.
func UpdateEndless(ctx app.Context, stage int) {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	s.Endless[stage]++
	ctx.LocalStorage().Set(localStorageScores, s)
	ctx.NewActionWithValue(EventScoreUpdate, s)
}

// IncrementEasyWin increments the win for an easy game.
func IncrementEasyWin(ctx app.Context) {
	incrementWin(ctx, Easy)
}

// IncrementMediumWin incremens the win for a medium game.
func IncrementMediumWin(ctx app.Context) {
	incrementWin(ctx, Medium)
}

// IncrementHardWin increments the win for a hard game.
func IncrementHardWin(ctx app.Context) {
	incrementWin(ctx, Hard)
}

func incrementWin(ctx app.Context, d Difficulty) {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	f := s.Basic[d]
	f.Win++
	s.Basic[d] = f
	ctx.LocalStorage().Set(localStorageScores, s)
	ctx.NewActionWithValue(EventScoreUpdate, s)
}

// LoadScores returns the currently stored Scores.
func LoadScores(ctx app.Context) Scores {
	s := newScores()
	ctx.LocalStorage().Get(localStorageScores, &s)
	return s
}
