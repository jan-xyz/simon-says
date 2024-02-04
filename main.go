package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	s := GenerateSequence(6)
	fmt.Printf("%v\n", s)
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
