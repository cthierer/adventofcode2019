package strategyguide

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cthierer/advent-of-code/rps"
)

type Entry struct {
	player1 rps.Action
	player2 rps.Action
}

func (e Entry) Player1() rps.Action {
	return e.player1
}

func (e Entry) Player2() rps.Action {
	return e.player2
}

var NoEntry = Entry{}

type StrategyGuide struct {
	started  bool
	position int
	entries  []Entry
}

func (s *StrategyGuide) Next() bool {
	if s.started {
		s.position += 1
	} else {
		s.started = true
	}

	return s.position < len(s.entries)
}

func (s *StrategyGuide) Get() Entry {
	if s.position >= len(s.entries) {
		return NoEntry
	}

	return s.entries[s.position]
}

func actionFromCol1(value string) (rps.Action, error) {
	switch value {
	case "A":
		return rps.Rock{}, nil
	case "B":
		return rps.Paper{}, nil
	case "C":
		return rps.Scissors{}, nil
	default:
		return nil, fmt.Errorf("invalid action identifier: %v", value)
	}
}

func actionFromCol2(value string) (rps.Action, error) {
	switch value {
	case "X":
		return rps.Rock{}, nil
	case "Y":
		return rps.Paper{}, nil
	case "Z":
		return rps.Scissors{}, nil
	default:
		return nil, fmt.Errorf("invalid action identifier: %v", value)
	}
}

func ParseStrategyGuide(input string) (*StrategyGuide, error) {
	var entries []Entry
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		cols := strings.Split(line, " ")
		if len(cols) != 2 {
			return nil, errors.New("unsupported format")
		}

		player1Action, err := actionFromCol1(cols[0])
		if err != nil {
			return nil, err
		}

		player2Action, err := actionFromCol2(cols[1])
		if err != nil {
			return nil, err
		}

		entries = append(entries, Entry{player1: player1Action, player2: player2Action})
	}

	return &StrategyGuide{entries: entries}, nil
}
