package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type events = string

const (
	click        events = "click"
	playSequence events = "playSequence"
	newGame      events = "newGame"
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
}

func (g *game) OnMount(ctx app.Context) {
	ctx.Handle(playSequence, g.playSequence)
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
	sequence, ok := a.Value.([]int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	fmt.Println("sequence:", sequence)
	for _, btnIndex := range sequence {
		btn := app.Window().GetElementByID(fmt.Sprintf("button%d", btnIndex))
		fmt.Println("found button:", btn.Truthy())
	}
}

func (g *game) handleNewGame(ctx app.Context, a app.Action) {
	fmt.Println("New Game")
	g.clicks = 0
	seq := GenerateSequence(4)
	g.sequence = seq
	ctx.NewActionWithValue(playSequence, seq)
}

func (g *game) handleClick(ctx app.Context, a app.Action) {
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

	fmt.Println("received click:", click)
	if g.sequence[g.clicks] != click {
		fmt.Println("YOU LOSE!")
		g.clicks = 0
		g.sequence = GenerateSequence(4)
		ctx.NewActionWithValue(playSequence, g.sequence)
		return
	}
	g.clicks++
	if len(g.sequence) == g.clicks {
		fmt.Println("YOU WIN!")
		g.clicks = 0
		g.sequence = GenerateSequence(4)
		ctx.NewActionWithValue(playSequence, g.sequence)
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
