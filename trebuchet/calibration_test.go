package trebuchet_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/trebuchet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLineFirstDigit(t *testing.T) {
	t.Parallel()

	t.Run("when the digit is at index 0", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "1abc2"}
		actual, err := line.FirstDigit()

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when the digit is not at index 0", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "pqr3stu8vwx"}
		actual, err := line.FirstDigit()

		require.Nil(t, err)
		assert.Equal(t, 3, actual)
	})

	t.Run("when the digit is at the last position", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "foo1"}
		actual, err := line.FirstDigit()

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when there are no digits", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "abcd"}
		_, err := line.FirstDigit()

		assert.ErrorIs(t, err, trebuchet.ErrNoDigitFound)
	})
}

func TestLineLastDigit(t *testing.T) {
	t.Parallel()

	t.Run("when the digit is at the last position", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "1abc2"}
		actual, err := line.LastDigit()

		require.Nil(t, err)
		assert.Equal(t, 2, actual)
	})

	t.Run("when the digit is not at the last position", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "pqr3stu8vwx"}
		actual, err := line.LastDigit()

		require.Nil(t, err)
		assert.Equal(t, 8, actual)
	})

	t.Run("when the last digit is at index 0", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "1foo"}
		actual, err := line.LastDigit()

		require.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("when there are no digits", func(t *testing.T) {
		t.Parallel()

		line := trebuchet.CalibrationLine{Value: "abcd"}
		_, err := line.LastDigit()

		assert.ErrorIs(t, err, trebuchet.ErrNoDigitFound)
	})
}

func TestLineCalibrationValue(t *testing.T) {
	lines := []struct {
		value    string
		expected int
	}{
		{
			value:    "1abc2",
			expected: 12,
		},
		{
			value:    "pqr3stu8vwx",
			expected: 38,
		},
		{
			value:    "a1b2c3d4e5f",
			expected: 15,
		},
		{
			value:    "treb7uchet",
			expected: 77,
		},
	}

	for idx, scenario := range lines {
		line := trebuchet.CalibrationLine{Value: scenario.value}
		actual, err := line.CalibrationValue()

		require.Nil(t, err, "error occurred at index %v", idx)
		assert.Equal(t, scenario.expected, actual, "scenario at index %v did not produce expected result", idx)
	}
}

func TestSumCalibrationValues(t *testing.T) {
	document := trebuchet.ParseDocument("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")
	actual, err := trebuchet.SumCalibrationValues(document)

	require.Nil(t, err)
	assert.Equal(t, 142, actual)
}
