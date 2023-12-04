package trebuchet_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/trebuchet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalibrationLineCalibrationValue(t *testing.T) {
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
