package trebuchet

import (
	"fmt"
	"strings"

	"github.com/cthierer/advent-of-code/trebuchet/digit"
)

type CalibrationDocument struct {
	lines []CalibrationLine
}

func ParseDocument(input string) *CalibrationDocument {
	values := strings.Split(input, "\n")
	lines := make([]CalibrationLine, len(values))

	for idx, value := range values {
		lines[idx] = CalibrationLine{Value: value}
	}

	return &CalibrationDocument{lines}
}

func SumCalibrationValues(document *CalibrationDocument) (int, error) {
	sum := 0

	for idx, line := range document.lines {
		calibrationValue, err := line.CalibrationValue()
		if err != nil {
			return 0, fmt.Errorf("unable to determine calibration value for line %v: %w", idx+1, err)
		}

		sum += calibrationValue
	}

	return sum, nil
}

type CalibrationLine struct {
	Value string
}

func (line CalibrationLine) firstDigit() (int, error) {
	return digit.FindFirstOccurrence(line.Value)
}

func (line CalibrationLine) lastDigit() (int, error) {
	return digit.FindLastOccurrence(line.Value)
}

func (line CalibrationLine) CalibrationValue() (int, error) {
	firstDigit, err := line.firstDigit()
	if err != nil {
		return 0, fmt.Errorf("unable to determine first digit: %w", err)
	}

	lastDigit, err := line.lastDigit()
	if err != nil {
		return 0, fmt.Errorf("unable to determine last digit: %w", err)
	}

	return firstDigit*padding(lastDigit) + lastDigit, nil
}

// adopted from: https://stackoverflow.com/a/28029597
func padding(value int) int {
	padding := 10
	for padding < value {
		padding *= 10
	}
	return padding
}
