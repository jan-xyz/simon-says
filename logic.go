package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/jan-xyz/simon-says/ui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type (
	difficulty     int
	sequenceLength int
)

const (
	easy difficulty = iota
	medium
	hard
)

var LocalStorageScores = "scores"

type scores struct {
	Basic   map[ui.Difficulty]score
	Endless map[int]int
}

type score struct {
	Win  int
	Loss int
}

var difficulties = map[ui.Difficulty]int{
	ui.Easy:   4,
	ui.Medium: 8,
	ui.Hard:   12,
}

type gameState int

const (
	gameStateNoGame gameState = iota
	gameStateSimonSays
	gameStatePlayerSays
	gameStateLost
	gameStateWon
)

func NewLogic() *logic {
	return &logic{}
}

type logic struct {
	difficulty   ui.Difficulty
	sequence     []int64
	clicks       int
	stage        int
	state        gameState
	storageMutex sync.Mutex
}

func (g *logic) simonSays(ctx app.Context, sequence []int64) {
	go func() {
		<-time.After(200 * time.Millisecond)
		for _, btnIndex := range sequence {
			ctx.NewAction(fmt.Sprintf(ui.EventPlayButton, btnIndex))
			<-time.After(time.Second)
		}
		g.state = gameStatePlayerSays
		ctx.NewActionWithValue(ui.EventStateChange, "Repeat what Simon said...")
	}()
}

func (g *logic) handleNewGame(ctx app.Context, a app.Action) {
	d, ok := a.Value.(ui.Difficulty)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	g.difficulty = d
	g.clicks = 0
	g.sequence = []int64{NextNumber()}
	g.stage = 1
	g.state = gameStateSimonSays
	ctx.NewActionWithValue(ui.EventStateChange, "Simon says...")
	g.simonSays(ctx, g.sequence)
}

func (g *logic) handleClick(ctx app.Context, a app.Action) {
	if g.state != gameStatePlayerSays {
		return
	}
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

	if g.sequence[g.clicks] != click {
		g.lostGame(ctx)
		return
	}
	g.clicks++
	if g.difficulty != ui.Endless && difficulties[g.difficulty] == g.clicks {
		g.wonGame(ctx)
		return
	}
	if g.clicks == g.stage {
		g.nextRound(ctx)
	}
}

func (g *logic) lostGame(ctx app.Context) {
	g.storageMutex.Lock()
	defer g.storageMutex.Unlock()
	g.state = gameStateLost
	ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You Lost in %s mode in stage %d. Franzi has a highscore of 21.", g.difficulty, len(g.sequence)))

	// increment losses
	s := &scores{}
	ctx.LocalStorage().Get(LocalStorageScores, s)
	if g.difficulty != ui.Endless {
		if f, ok := s.Basic[g.difficulty]; ok {
			f.Loss++
			s.Basic[g.difficulty] = f
		} else {
			if s.Basic == nil {
				s.Basic = map[ui.Difficulty]score{}
			}
			s.Basic[g.difficulty] = score{Loss: 1}
		}
	} else {
		if s.Endless == nil {
			s.Endless = map[int]int{}
		}
		s.Endless[len(g.sequence)]++
	}
	ctx.LocalStorage().Set(LocalStorageScores, s)
}

func (g *logic) wonGame(ctx app.Context) {
	g.storageMutex.Lock()
	defer g.storageMutex.Unlock()
	g.state = gameStateWon
	ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You Won in %s mode. Start a New Game", g.difficulty))

	// increment wins
	s := &scores{}
	ctx.LocalStorage().Get(LocalStorageScores, s)
	if f, ok := s.Basic[g.difficulty]; ok {
		f.Win++
		s.Basic[g.difficulty] = f
	} else {
		if s.Basic == nil {
			s.Basic = map[ui.Difficulty]score{}
		}
		s.Basic[g.difficulty] = score{Win: 1}
	}
	ctx.LocalStorage().Set(LocalStorageScores, s)
}

func (g *logic) nextRound(ctx app.Context) {
	g.clicks = 0
	g.stage++
	g.state = gameStateSimonSays
	g.sequence = append(g.sequence, NextNumber())
	ctx.NewActionWithValue(ui.EventStateChange, "Simon says...")
	ctx.After(1*time.Second, func(ctx app.Context) {
		g.simonSays(ctx, g.sequence)
	})
}

func NextNumber() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
