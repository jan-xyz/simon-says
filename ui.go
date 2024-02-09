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
	click      events = "click"
	simonSays  events = "playSequence"
	playButton events = "play%d"
	newGame    events = "newGame"
)

type gameState int

const (
	gameStateNoGame gameState = iota
	gameStateSimonSays
	gameStatePlayerSays
	gameStateLost
	gameStateWon
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
	State    gameState
}

func (g *game) OnMount(ctx app.Context) {
	// TODO: move logic out of the UI thread
	ctx.Handle(simonSays, g.simonSays)
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

	txt := ""
	switch g.State {
	case gameStateNoGame:
		txt = "Start a New Game"
	case gameStatePlayerSays:
		txt = "Your turn"
	case gameStateSimonSays:
		txt = "Simon says..."
	case gameStateLost:
		txt = "You Lost. Start a New Game"
	case gameStateWon:
		txt = "You Won. Start a New Game"
	}

	// TODO: styling
	gameStateText := app.Div().Class("game-state").Text(txt)
	newGameButton := app.Button().Class("new-game").Text("New Game").OnClick(func(ctx app.Context, _ app.Event) {
		ctx.NewAction(newGame)
	})
	return app.Div().Class("fill", "background").Body(
		gameField,
		gameStateText,
		newGameButton,
	)
}

func (g *game) simonSays(ctx app.Context, a app.Action) {
	g.Update()
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
	})
	g.State = gameStatePlayerSays
	g.Update()
}

func (g *game) handleNewGame(ctx app.Context, a app.Action) {
	fmt.Println("New Game")
	g.clicks = 0
	// TODO: allow setting difficulty
	g.sequence = GenerateSequence(4)
	g.stage = 1
	g.State = gameStateSimonSays
	ctx.NewActionWithValue(simonSays, g.sequence[:1])
}

func (g *game) handleClick(ctx app.Context, a app.Action) {
	if g.State != gameStatePlayerSays {
		fmt.Println("no game")
		return
	}
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

	fmt.Println("received click:", click)
	if g.sequence[g.clicks] != click {
		g.State = gameStateLost
		return
	}
	g.clicks++
	if len(g.sequence) == g.clicks {
		g.State = gameStateWon
		return
	}
	if g.clicks == g.stage {
		g.clicks = 0
		g.stage++
		g.State = gameStateSimonSays
		ctx.After(1*time.Second, func(ctx app.Context) {
			ctx.NewActionWithValue(simonSays, g.sequence[:g.stage])
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
