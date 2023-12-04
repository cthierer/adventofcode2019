package cubegame

import (
	"fmt"
	"strings"
)

type GameSet struct {
	Games []*Game
}

func (set *GameSet) String() string {
	games := make([]string, len(set.Games))
	for idx, game := range set.Games {
		games[idx] = game.String()
	}

	return strings.Join(games, "\n")
}

func ParseGameSet(lines string) (*GameSet, error) {
	values := strings.Split(lines, "\n")

	parsed := GameSet{}
	parsed.Games = make([]*Game, len(values))

	var err error

	for idx, line := range values {
		parsed.Games[idx], err = ParseGame(line)
		if err != nil {
			return nil, fmt.Errorf("unable to parse game at line %v: %w", idx+1, err)
		}
	}

	return &parsed, nil
}

func QueryPossibleGames(games *GameSet, query *CubeSet) *GameSet {
	possible := GameSet{}
	colors := query.Colors()

	for _, game := range games.Games {
		match := true

		for _, color := range colors {
			if game.GetMaxCount(color) > query.GetCount(color) {
				match = false
				break
			}
		}

		if match {
			possible.Games = append(possible.Games, game)
		}
	}

	return &possible
}
