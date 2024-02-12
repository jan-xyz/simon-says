package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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
	difficulty ui.Difficulty
	sequence   []int64
	clicks     int
	stage      int
	state      gameState
}

func (g *logic) simonSays(ctx app.Context, a app.Action) {
	sequence, ok := a.Value.([]int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

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
	ctx.NewActionWithValue(ui.EventSimonSays, g.sequence)
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
		g.state = gameStateLost
		ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You Lost in %s mode with a streak of %d. Franzi has a highscore of 21.", g.difficulty, len(g.sequence)))
		return
	}
	g.clicks++
	if g.difficulty != ui.Endless && difficulties[g.difficulty] == g.clicks {
		g.state = gameStateWon
		ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You Won in %s mode. Start a New Game", g.difficulty))
		return
	}
	if g.clicks == g.stage {
		g.clicks = 0
		g.stage++
		g.state = gameStateSimonSays
		g.sequence = append(g.sequence, NextNumber())
		ctx.NewActionWithValue(ui.EventStateChange, "Simon says...")
		ctx.After(1*time.Second, func(ctx app.Context) {
			ctx.NewActionWithValue(ui.EventSimonSays, g.sequence)
		})
	}
}

func NextNumber() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
