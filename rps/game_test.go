package rps_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/rps"
)

type round struct {
	player1 rps.Action
	player2 rps.Action
}

func newRound(player1, player2 rps.Action) round {
	return round{player1, player2}
}

func (r round) Player1() rps.Action {
	return r.player1
}

func (r round) Player2() rps.Action {
	return r.player2
}

func TestGame(t *testing.T) {
	g := rps.Game{}

	g.Play(newRound(rps.Rock{}, rps.Paper{}))
	g.Play(newRound(rps.Paper{}, rps.Rock{}))
	g.Play(newRound(rps.Scissors{}, rps.Scissors{}))

	_, myScore := g.Scores()
	if myScore != 15 {
		t.Fail()
	}
}
