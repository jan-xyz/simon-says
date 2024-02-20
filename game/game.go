// Package game is the main game logic. It implements the simon says parts
// and the round/stage tracking.
package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/jan-xyz/simon-says/storage"
	"github.com/jan-xyz/simon-says/ui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var difficulties = map[storage.Difficulty]int{
	storage.Easy:   4,
	storage.Medium: 8,
	storage.Hard:   12,
}

var speeds = map[storage.Speed]time.Duration{
	storage.Slow:   time.Second,
	storage.Normal: 800 * time.Millisecond,
	storage.Fast:   400 * time.Millisecond,
}

type gameState int

const (
	gameStateNoGame gameState = iota
	gameStateSimonSays
	gameStatePlayerSays
	gameStateLost
	gameStateWon
)

// New is the factory for the Game component.
func New() *Game {
	return &Game{}
}

// Game is the main Game component of the game. Create a new instance by
type Game struct {
	difficulty   storage.Difficulty
	speed        storage.Speed
	sequence     []int64
	clicks       int
	stage        int
	state        gameState
	storageMutex sync.Mutex
}

func (g *Game) simonSays(ctx app.Context, sequence []int64) {
	ctx.Async(func() {
		<-time.After(200 * time.Millisecond)
		for _, btnIndex := range sequence {
			eventName := fmt.Sprintf(ui.EventPlayButton, btnIndex)
			ctx.NewActionWithValue(eventName, speeds[g.speed]-100*time.Millisecond)
			<-time.After(speeds[g.speed])
		}
		g.state = gameStatePlayerSays
		ctx.NewActionWithValue(ui.EventStateChange, "Repeat what Simon said...")
	})
}

// HandleNewGame is the handler to invoke to start a new game.
func (g *Game) HandleNewGame(ctx app.Context, a app.Action) {
	d, ok := a.Value.(ui.NewGameSettings)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	g.difficulty = d.Difficulty
	g.speed = d.Speed
	g.clicks = 0
	g.sequence = []int64{nextNumber()}
	g.stage = 1
	g.state = gameStateSimonSays
	ctx.NewActionWithValue(ui.EventStateChange, "Simon says...")
	g.simonSays(ctx, g.sequence)
}

// HandleClick is the handler to invoke on user input when it's the user's turn.
func (g *Game) HandleClick(ctx app.Context, a app.Action) {
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
	if g.difficulty != storage.Endless && difficulties[g.difficulty] == g.clicks {
		g.wonGame(ctx)
		return
	}
	if g.clicks == g.stage {
		g.nextRound(ctx)
	}
}

func (g *Game) lostGame(ctx app.Context) {
	g.storageMutex.Lock()
	defer g.storageMutex.Unlock()
	g.state = gameStateLost
	ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You lost in %s mode in stage %d. Franzi has a highscore of 21.", g.difficulty, len(g.sequence)))

	// increment losses
	switch g.difficulty {
	case storage.Easy:
		storage.IncrementEasyLoss(ctx)
	case storage.Medium:
		storage.IncrementMediumLoss(ctx)
	case storage.Hard:
		storage.IncrementHardLoss(ctx)
	case storage.Endless:
		storage.UpdateEndless(ctx, len(g.sequence))
	default:
		app.Log("Difficulty not supported")
	}
}

func (g *Game) wonGame(ctx app.Context) {
	g.storageMutex.Lock()
	defer g.storageMutex.Unlock()
	g.state = gameStateWon
	ctx.NewActionWithValue(ui.EventStateChange, fmt.Sprintf("You won in %s mode. Start a New Game", g.difficulty))

	// increment wins
	switch g.difficulty {
	case storage.Easy:
		storage.IncrementEasyWin(ctx)
	case storage.Medium:
		storage.IncrementMediumWin(ctx)
	case storage.Hard:
		storage.IncrementHardWin(ctx)
	default:
		app.Log("Difficulty not supported")
	}
}

func (g *Game) nextRound(ctx app.Context) {
	g.clicks = 0
	g.stage++
	g.state = gameStateSimonSays
	g.sequence = append(g.sequence, nextNumber())
	ctx.NewActionWithValue(ui.EventStateChange, "Simon says...")
	ctx.After(1*time.Second, func(ctx app.Context) {
		g.simonSays(ctx, g.sequence)
	})
}

func nextNumber() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(4))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
