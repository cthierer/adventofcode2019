package digit_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/trebuchet/digit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindFirstOccurrence(t *testing.T) {
	t.Parallel()

	t.Run("when the digit is at index 0", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindFirstOccurrence("1abc2")

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when the digit is not at index 0", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindFirstOccurrence("pqr3stu8vwx")

		require.Nil(t, err)
		assert.Equal(t, 3, actual)
	})

	t.Run("when the digit is at the last position", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindFirstOccurrence("foo1")

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when the digit is spelled out with letters", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindFirstOccurrence("two1nine")

		require.Nil(t, err)
		assert.Equal(t, 2, actual)
	})

	t.Run("recognizes all possible digits", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			value    string
			expected int
		}{
			{
				value:    "0",
				expected: 0,
			},
			{
				value:    "1",
				expected: 1,
			},
			{
				value:    "one",
				expected: 1,
			},
			{
				value:    "2",
				expected: 2,
			},
			{
				value:    "two",
				expected: 2,
			},
			{
				value:    "3",
				expected: 3,
			},
			{
				value:    "three",
				expected: 3,
			},
			{
				value:    "4",
				expected: 4,
			},
			{
				value:    "four",
				expected: 4,
			},
			{
				value:    "5",
				expected: 5,
			},
			{
				value:    "five",
				expected: 5,
			},
			{
				value:    "6",
				expected: 6,
			},
			{
				value:    "six",
				expected: 6,
			},
			{
				value:    "7",
				expected: 7,
			},
			{
				value:    "seven",
				expected: 7,
			},
			{
				value:    "8",
				expected: 8,
			},
			{
				value:    "eight",
				expected: 8,
			},
			{
				value:    "9",
				expected: 9,
			},
			{
				value:    "nine",
				expected: 9,
			},
		}

		for idx, scenario := range tests {
			actual, err := digit.FindFirstOccurrence(scenario.value)

			require.Nil(t, err, "error occurred at index %v", idx)
			assert.Equal(t, scenario.expected, actual, "scenario at index %v did not produce expected result", idx)
		}
	})

	t.Run("when there are no digits", func(t *testing.T) {
		t.Parallel()

		_, err := digit.FindFirstOccurrence("abcd")

		assert.ErrorIs(t, err, digit.ErrNotFound)
	})
}

func TestFindLastOccurrence(t *testing.T) {
	t.Parallel()

	t.Run("when the digit is at the last position", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindLastOccurrence("1abc2")

		require.Nil(t, err)
		assert.Equal(t, 2, actual)
	})

	t.Run("when the digit is not at the last position", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindLastOccurrence("pqr3stu8vwx")

		require.Nil(t, err)
		assert.Equal(t, 8, actual)
	})

	t.Run("when the last digit is at index 0", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindLastOccurrence("1foo")

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when the digit is spelled out with letters", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindLastOccurrence("two1nine")

		require.Nil(t, err)
		assert.Equal(t, 9, actual)
	})

	t.Run("when the digit is duplicated multiple times", func(t *testing.T) {
		t.Parallel()

		actual, err := digit.FindLastOccurrence("twofourninefour")

		require.Nil(t, err)
		assert.Equal(t, 4, actual)
	})

	t.Run("recognizes all possible digits", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			value    string
			expected int
		}{
			{
				value:    "0",
				expected: 0,
			},
			{
				value:    "1",
				expected: 1,
			},
			{
				value:    "one",
				expected: 1,
			},
			{
				value:    "2",
				expected: 2,
			},
			{
				value:    "two",
				expected: 2,
			},
			{
				value:    "3",
				expected: 3,
			},
			{
				value:    "three",
				expected: 3,
			},
			{
				value:    "4",
				expected: 4,
			},
			{
				value:    "four",
				expected: 4,
			},
			{
				value:    "5",
				expected: 5,
			},
			{
				value:    "five",
				expected: 5,
			},
			{
				value:    "6",
				expected: 6,
			},
			{
				value:    "six",
				expected: 6,
			},
			{
				value:    "7",
				expected: 7,
			},
			{
				value:    "seven",
				expected: 7,
			},
			{
				value:    "8",
				expected: 8,
			},
			{
				value:    "eight",
				expected: 8,
			},
			{
				value:    "9",
				expected: 9,
			},
			{
				value:    "nine",
				expected: 9,
			},
		}

		for idx, scenario := range tests {
			actual, err := digit.FindLastOccurrence(scenario.value)

			require.Nil(t, err, "error occurred at index %v", idx)
			assert.Equal(t, scenario.expected, actual, "scenario at index %v did not produce expected result", idx)
		}
	})

	t.Run("when there are no digits", func(t *testing.T) {
		t.Parallel()

		_, err := digit.FindLastOccurrence("abcd")

		assert.ErrorIs(t, err, digit.ErrNotFound)
	})
}
