package strategyguide_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/rps"
	"github.com/cthierer/advent-of-code/strategyguide"
)

const sampleInput = `A Y
B X
C Z`

func isRock(action rps.Action) bool {
	switch action.(type) {
	case rps.Rock:
		return true
	default:
		return false
	}
}

func isPaper(action rps.Action) bool {
	switch action.(type) {
	case rps.Paper:
		return true
	default:
		return false
	}
}

func isScissors(action rps.Action) bool {
	switch action.(type) {
	case rps.Scissors:
		return true
	default:
		return false
	}
}

func TestParseStrategyGuide(t *testing.T) {
	guide, err := strategyguide.ParseStrategyGuide(sampleInput)
	if err != nil {
		t.Logf("parsing guide failed: %v", err)
		t.FailNow()
	}

	if guide.Next() != true {
		t.Fail()
	}

	row1 := guide.Get()
	if row1 == strategyguide.NoEntry || !isRock(row1.Player1()) || !isRock(row1.Player2()) {
		t.Fail()
	}

	if guide.Next() != true {
		t.Fail()
	}

	row2 := guide.Get()
	if row2 == strategyguide.NoEntry || !isPaper(row2.Player1()) || !isRock(row2.Player2()) {
		t.Fail()
	}

	if guide.Next() != true {
		t.Fail()
	}

	row3 := guide.Get()
	if row3 == strategyguide.NoEntry || !isScissors(row3.Player1()) || !isRock(row3.Player2()) {
		t.Fail()
	}
}
