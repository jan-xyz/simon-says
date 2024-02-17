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
	difficulty   storage.Difficulty
	sequence     []int64
	clicks       int
	stage        int
	state        gameState
	storageMutex sync.Mutex
}

func (g *logic) simonSays(ctx app.Context, sequence []int64) {
	fmt.Println(sequence)
	ctx.Async(func() {
		<-time.After(200 * time.Millisecond)
		for _, btnIndex := range sequence {
			ctx.NewAction(fmt.Sprintf(ui.EventPlayButton, btnIndex))
			<-time.After(time.Second)
		}
		g.state = gameStatePlayerSays
		ctx.NewActionWithValue(ui.EventStateChange, "Repeat what Simon said...")
	})
}

func (g *logic) HandleNewGame(ctx app.Context, a app.Action) {
	d, ok := a.Value.(storage.Difficulty)
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

func (g *logic) HandleClick(ctx app.Context, a app.Action) {
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

func (g *logic) lostGame(ctx app.Context) {
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

func (g *logic) wonGame(ctx app.Context) {
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
