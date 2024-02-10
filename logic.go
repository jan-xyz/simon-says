package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

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

var difficulties = map[difficulty]sequenceLength{
	easy:   4,
	medium: 8,
	hard:   12,
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
	return &logic{
		sequence: []int64{},
	}
}

type logic struct {
	app.Compo

	sequence []int64
	clicks   int
	stage    int
	state    gameState
}

func (g *logic) simonSays(ctx app.Context, a app.Action) {
	sequence, ok := a.Value.([]int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	fmt.Println("sequence:", sequence)

	go func() {
		<-time.After(200 * time.Millisecond)
		for _, btnIndex := range sequence {
			fmt.Println("sending", btnIndex)
			ctx.NewAction(fmt.Sprintf(eventPlayButton, btnIndex))
			<-time.After(time.Second)
		}
		g.state = gameStatePlayerSays
		ctx.NewActionWithValue(eventStateChange, g.state)
	}()
}

func (g *logic) handleNewGame(ctx app.Context, a app.Action) {
	d, ok := a.Value.(difficulty)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	g.clicks = 0
	g.sequence = GenerateSequence(difficulties[d])
	g.stage = 1
	g.state = gameStateSimonSays
	ctx.NewActionWithValue(eventStateChange, g.state)
	ctx.NewActionWithValue(eventSimonSays, g.sequence[:1])
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

	fmt.Println("received click:", click)
	if g.sequence[g.clicks] != click {
		g.state = gameStateLost
		ctx.NewActionWithValue(eventStateChange, g.state)
		return
	}
	g.clicks++
	if len(g.sequence) == g.clicks {
		g.state = gameStateWon
		ctx.NewActionWithValue(eventStateChange, g.state)
		return
	}
	if g.clicks == g.stage {
		g.clicks = 0
		g.stage++
		g.state = gameStateSimonSays
		ctx.NewActionWithValue(eventStateChange, g.state)
		ctx.After(1*time.Second, func(ctx app.Context) {
			ctx.NewActionWithValue(eventSimonSays, g.sequence[:g.stage])
		})
	}
}

func GenerateSequence(l sequenceLength) []int64 {
	seq := []int64{}
	for i := 0; i < int(l); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(4))
		if err != nil {
			panic(err)
		}
		seq = append(seq, n.Int64())

	}
	return seq
}
