package section

import (
	"errors"
	"strings"
)

type Pair struct {
	assignment1 Range
	assignment2 Range
}

func NewPair(assignment1, assignment2 Range) Pair {
	return Pair{assignment1, assignment2}
}

func ParsePair(value string) (Pair, error) {
	ranges := strings.Split(value, ",")
	if len(ranges) < 2 {
		return Pair{}, errors.New("unexpected pair format")
	}

	assignment1, err := ParseRange(ranges[0])
	if err != nil {
		return Pair{}, err
	}

	assignment2, err := ParseRange(ranges[1])
	if err != nil {
		return Pair{}, err
	}

	return NewPair(assignment1, assignment2), nil
}

func (p Pair) Assignment1() Range {
	return p.assignment1
}

func (p Pair) Assignment2() Range {
	return p.assignment2
}

func (p Pair) Overlapping() bool {
	return p.assignment1.Contains(p.assignment2) || p.assignment2.Contains(p.assignment1)
}
