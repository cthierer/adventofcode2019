package trebuchet

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrNoDigitFound = errors.New("no digit found")

type CalibrationLine struct {
	Value string
}

func (line CalibrationLine) FirstDigit() (int, error) {
	var firstDigit int64
	var err error

	for _, character := range line.Value {
		firstDigit, err = strconv.ParseInt(string(character), 10, 64)
		if err == nil {
			return int(firstDigit), nil
		}
	}

	return 0, ErrNoDigitFound
}

func (line CalibrationLine) LastDigit() (int, error) {
	var lastDigit int64
	var err error

	for i := len(line.Value) - 1; i >= 0; i-- {
		character := line.Value[i]
		lastDigit, err = strconv.ParseInt(string(character), 10, 64)
		if err == nil {
			return int(lastDigit), nil
		}
	}

	return 0, ErrNoDigitFound
}

func (line CalibrationLine) CalibrationValue() (int, error) {
	firstDigit, err := line.FirstDigit()
	if err != nil {
		return 0, fmt.Errorf("unable to determine first digit: %w", err)
	}

	lastDigit, err := line.LastDigit()
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

	for _, line := range document.lines {
		calibrationValue, err := line.CalibrationValue()
		if err != nil {
			return 0, fmt.Errorf("unable to determine calibration value: %w", err)
		}

		sum += calibrationValue
	}

	return sum, nil
}
