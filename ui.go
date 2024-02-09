package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	click        events = "click"
	playSequence events = "playSequence"
	playButton   events = "play%d"
	newGame      events = "newGame"
)

type gameState int

const (
	gameStateNoGame gameState = iota
	gameStateSimonSays
	gameStatePlayerSays
)

func NewGame() *game {
	return &game{
		sequence: []int64{},
	}
}

type game struct {
	app.Compo

	sequence []int64
	clicks   int
	stage    int
	state    gameState
}

func (g *game) OnMount(ctx app.Context) {
	ctx.Handle(playSequence, g.playSequence)
	ctx.Handle(click, g.handleClick)
	ctx.Handle(newGame, g.handleNewGame)
}

func (g *game) Render() app.UI {
	gameField := app.Div().Class("game-field")

	firstButton := NewButton(0)
	secondButton := NewButton(1)
	thirdButton := NewButton(2)
	fourthButton := NewButton(3)

	gameField.Body(
		firstButton,
		secondButton,
		thirdButton,
		fourthButton,
	)

	return app.Div().Class("fill", "background").Body(
		gameField,
		app.Button().Text("New Game").OnClick(func(ctx app.Context, _ app.Event) {
			ctx.NewAction(newGame)
		}),
	)
}

func (g *game) playSequence(ctx app.Context, a app.Action) {
	g.state = gameStateSimonSays
	sequence, ok := a.Value.([]int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	fmt.Println("sequence:", sequence)

	ctx.Async(func() {
		// TODO: This is so weird. I somehow need to wait before I can do the next
		//       action, otherwise it just won't update the DOM.
		<-time.After(200 * time.Millisecond)
		for _, btnIndex := range sequence {
			fmt.Println("sending", btnIndex)
			ctx.NewAction(fmt.Sprintf(playButton, btnIndex))
			<-time.After(time.Second)
		}
		g.state = gameStatePlayerSays
	})
}

func (g *game) handleNewGame(ctx app.Context, a app.Action) {
	fmt.Println("New Game")
	g.clicks = 0
	g.sequence = GenerateSequence(4)
	g.stage = 1
	ctx.NewActionWithValue(playSequence, g.sequence[:1])
}

func (g *game) handleClick(ctx app.Context, a app.Action) {
	if g.state != gameStatePlayerSays {
		return
	}
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	if len(g.sequence) <= g.clicks {
		fmt.Println("game is over")
		return
	}

	fmt.Println("received click:", click)
	if g.sequence[g.clicks] != click {
		fmt.Println("YOU LOSE!")
		return
	}
	g.clicks++
	if len(g.sequence) == g.clicks {
		fmt.Println("YOU WIN!")
		return
	}
	if g.clicks == g.stage {
		g.clicks = 0
		g.stage++
		ctx.After(300*time.Millisecond, func(ctx app.Context) {
			ctx.NewActionWithValue(playSequence, g.sequence[:g.stage])
		})
	}
}

func GenerateSequence(l int) []int64 {
	seq := []int64{}
	for i := 0; i < l; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(4))
		if err != nil {
			panic(err)
		}
		seq = append(seq, n.Int64())

	}
	return seq
}
