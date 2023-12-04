package cubegame_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cubegame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGameSet(t *testing.T) {
	set, err := cubegame.ParseGameSet(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

	require.Nil(t, err)
	assert.Len(t, set.Games, 5)
}

func TestQueryPossibleGames(t *testing.T) {
	games, err := cubegame.ParseGameSet(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)
	require.Nil(t, err)

	query := &cubegame.CubeSet{}
	query.SetCount(cubegame.Red, 12)
	query.SetCount(cubegame.Green, 13)
	query.SetCount(cubegame.Blue, 14)

	possible := cubegame.QueryPossibleGames(games, query)
	assert.Len(t, possible.Games, 3)
	assert.Equal(t, 1, possible.Games[0].ID)
	assert.Equal(t, 2, possible.Games[1].ID)
	assert.Equal(t, 5, possible.Games[2].ID)
}
