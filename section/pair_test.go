package section_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/section"
)

func TestParsePair(t *testing.T) {
	p, err := section.ParsePair("2-4,6-8")
	if err != nil {
		t.FailNow()
	}

	if p.Assignment1().Start() != 2 || p.Assignment1().Stop() != 4 {
		t.Fail()
	}

	if p.Assignment2().Start() != 6 || p.Assignment2().Stop() != 8 {
		t.Fail()
	}
}

func TestPairOverlapping(t *testing.T) {
	scenarios := []struct {
		pair     section.Pair
		expected bool
	}{
		{
			pair:     section.NewPair(section.NewRange(2, 4), section.NewRange(6, 8)),
			expected: false,
		},
		{
			pair:     section.NewPair(section.NewRange(2, 3), section.NewRange(4, 5)),
			expected: false,
		},
		{
			pair:     section.NewPair(section.NewRange(5, 7), section.NewRange(7, 9)),
			expected: false,
		},
		{
			pair:     section.NewPair(section.NewRange(2, 8), section.NewRange(3, 7)),
			expected: true,
		},
		{
			pair:     section.NewPair(section.NewRange(6, 6), section.NewRange(4, 6)),
			expected: true,
		},
		{
			pair:     section.NewPair(section.NewRange(2, 6), section.NewRange(4, 8)),
			expected: false,
		},
	}

	for _, s := range scenarios {
		if s.pair.Overlapping() != s.expected {
			t.Fail()
		}
	}
}
