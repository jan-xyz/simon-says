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
)

func NewGame() *game {
	return &game{
		buttonIds: []string{
			"firstButton",
			"secondButton",
			"thirdButton",
			"fourthButton",
		},
		sequence: GenerateSequence(4),
	}
}

type game struct {
	app.Compo

	buttonIds []string
	sequence  []int64
	clicks    int
}

func (h *game) OnMount(ctx app.Context) {
	ctx.Handle(playSequence, h.playSequence)
}

func (h *game) Render() app.UI {
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
		app.Button().Text("New Game").OnClick(h.handleNewGame),
	)
}

func (h *game) playSequence(ctx app.Context, a app.Action) {
	sequence, ok := a.Value.([]int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	fmt.Println("sequence:", sequence)
	for _, btnIndex := range sequence {
		btn := app.Window().GetElementByID(h.buttonIds[btnIndex])
		fmt.Println(btn)
		ctx.After(time.Second, func(ctx app.Context) {
		})
	}
}

func (s *game) handleNewGame(ctx app.Context, a app.Event) {
	fmt.Println("New Game")
}

func (s *game) handleClick(_ app.Context, a app.Action) {
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

	fmt.Println("received click:", click)
	if s.sequence[s.clicks] != click {
		fmt.Println("YOU LOSE!")
		s.clicks = 0
		s.sequence = GenerateSequence(4)
		return
	}
	s.clicks++
	if len(s.sequence) == s.clicks {
		fmt.Println("YOU WIN!")
		s.clicks = 0
		s.sequence = GenerateSequence(4)
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
	fmt.Println(seq)
	return seq
}
