package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"slices"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func NewSimonSaysLogic() *simonSaysLogic {
	return &simonSaysLogic{
		sequence: GenerateSequence(4),
		clicks:   []int64{},
	}
}

type simonSaysLogic struct {
	sequence []int64
	clicks   []int64
}

func (s *simonSaysLogic) handleClick(_ app.Context, a app.Action) {
	click, ok := a.Value.(int64)
	if !ok {
		fmt.Println("wrong type")
		return
	}

	fmt.Println("clicked button:", click)
	s.clicks = append(s.clicks, click)
	if len(s.clicks) == 4 {
		if slices.Equal(s.sequence, s.clicks) {
			fmt.Println("YOU WIN!")
		} else {
			fmt.Println("YOU LOSE!")
		}
		fmt.Println("RESTART")
		s.clicks = []int64{}
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
	fmt.Println("sequence:", seq)
	return seq
}
