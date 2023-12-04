package cubegame_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cubegame"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseColor(t *testing.T) {
	t.Parallel()

	t.Run("parses \"blue\" as blue", func(t *testing.T) {
		t.Parallel()

		actual, err := cubegame.ParseColor("blue")

		require.Nil(t, err)
		assert.Equal(t, cubegame.Blue, actual)
	})

	t.Run("parses \"red\" as red", func(t *testing.T) {
		t.Parallel()

		actual, err := cubegame.ParseColor("red")

		require.Nil(t, err)
		assert.Equal(t, cubegame.Red, actual)
	})

	t.Run("parses \"green\" as green", func(t *testing.T) {
		t.Parallel()

		actual, err := cubegame.ParseColor("green")

		require.Nil(t, err)
		assert.Equal(t, cubegame.Green, actual)
	})

	t.Run("fails on an unknown color", func(t *testing.T) {
		t.Parallel()

		_, err := cubegame.ParseColor("foo")

		assert.ErrorIs(t, err, cubegame.ErrUnknownColor)
	})
}
