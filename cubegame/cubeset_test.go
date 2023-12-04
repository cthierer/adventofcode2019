package cubegame_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cubegame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCubeSet(t *testing.T) {
	set, err := cubegame.ParseCubeSet("3 green, 4 blue, 1 red")

	require.Nil(t, err)

	green := set.GetCount(cubegame.Green)
	assert.Equal(t, 3, green)

	blue := set.GetCount(cubegame.Blue)
	assert.Equal(t, 4, blue)

	red := set.GetCount(cubegame.Red)
	assert.Equal(t, 1, red)
}
