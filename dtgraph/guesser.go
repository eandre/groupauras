package dtgraph

import "time"

type Guesser interface {
	Guess(dt time.Duration) int
}

type BigWigsGuesser struct{}

func (g *BigWigsGuesser) Guess(dt time.Duration) int {
	return 0
}
