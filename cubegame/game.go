package cubegame

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidGameFormat = errors.New("line is malformed")

type Game struct {
	ID    int
	Turns []*CubeSet
}

func (game *Game) GetMaxCount(color Color) int {
	max := 0
	for _, turn := range game.Turns {
		count := turn.GetCount(color)
		if count > max {
			max = count
		}
	}
	return max
}

func (game *Game) String() string {
	turns := make([]string, len(game.Turns))
	for idx, turn := range game.Turns {
		turns[idx] = turn.String()
	}

	return fmt.Sprintf("Game %v: %v", game.ID, strings.Join(turns, "; "))
}

func ParseGame(line string) (*Game, error) {
	trimmed := strings.TrimSpace(line)

	parts := strings.SplitN(trimmed, ":", 2)
	if len(parts) < 2 {
		return nil, ErrInvalidGameFormat
	}

	var err error
	parsed := Game{}

	parsed.ID, err = parseGameID(parts[0])
	if err != nil {
		return nil, fmt.Errorf("unable to parse game ID: %w", err)
	}

	turns := strings.Split(parts[1], ";")
	parsed.Turns = make([]*CubeSet, len(turns))
	for idx, turn := range turns {
		parsed.Turns[idx], err = ParseCubeSet(turn)
		if err != nil {
			return nil, fmt.Errorf("unable to parse turn at index %v: %w", idx, err)
		}
	}

	return &parsed, nil
}

func parseGameID(label string) (int, error) {
	trimmed := strings.TrimSpace(label)

	parts := strings.SplitN(trimmed, " ", 2)
	if len(parts) < 2 {
		return 0, ErrInvalidGameFormat
	}

	return strconv.Atoi(parts[1])
}
