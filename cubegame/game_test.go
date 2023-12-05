package cubegame_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cubegame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGame(t *testing.T) {
	game, err := cubegame.ParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	require.Nil(t, err)

	assert.Equal(t, 1, game.ID)
	assert.Len(t, game.Turns, 3)
}

func TestGameGetMaxCount(t *testing.T) {
	game, err := cubegame.ParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	require.Nil(t, err)

	red := game.MaxCount(cubegame.Red)
	assert.Equal(t, 4, red)

	green := game.MaxCount(cubegame.Green)
	assert.Equal(t, 2, green)

	blue := game.MaxCount(cubegame.Blue)
	assert.Equal(t, 6, blue)
}
