package section_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/section"
)

func TestParseRange(t *testing.T) {
	r, err := section.ParseRange("2-4")
	if err != nil {
		t.FailNow()
	}

	if r.Start() != 2 || r.Stop() != 4 {
		t.Fail()
	}
}

func TestRangeContains(t *testing.T) {
	scenarios := []struct {
		range1   section.Range
		range2   section.Range
		expected int
	}{
		{
			range1:   section.NewRange(2, 4),
			range2:   section.NewRange(6, 8),
			expected: 0,
		},
		{
			range1:   section.NewRange(2, 3),
			range2:   section.NewRange(4, 5),
			expected: 0,
		},
		{
			range1:   section.NewRange(5, 7),
			range2:   section.NewRange(7, 9),
			expected: 0,
		},
		{
			range1:   section.NewRange(2, 8),
			range2:   section.NewRange(3, 7),
			expected: 1,
		},
		{
			range1:   section.NewRange(6, 6),
			range2:   section.NewRange(4, 6),
			expected: 2,
		},
		{
			range1:   section.NewRange(2, 6),
			range2:   section.NewRange(4, 8),
			expected: 0,
		},
	}

	for _, s := range scenarios {
		overlapping1 := s.range1.Contains(s.range2)
		overlapping2 := s.range2.Contains(s.range1)

		if s.expected == 0 && (overlapping1 || overlapping2) {
			t.Logf("expected no overlap, but got: %v, %v", overlapping1, overlapping2)
			t.Fail()
		}

		if s.expected == 1 && (!overlapping1 || overlapping2) {
			t.Logf("expected 1 to contain 2, but got: %v, %v", overlapping1, overlapping2)
			t.Fail()
		}

		if s.expected == 2 && (overlapping1 || !overlapping2) {
			t.Logf("expected 2 to contain 1, but got: %v, %v", overlapping1, overlapping2)
			t.Fail()
		}

		if !s.range1.Contains(s.range1) || !s.range2.Contains(s.range2) {
			t.Log("expected each set to contain itself")
			t.Fail()
		}
	}
}

func TestRangeOverlaps(t *testing.T) {
	scenarios := []struct {
		range1   section.Range
		range2   section.Range
		expected bool
	}{
		{
			range1:   section.NewRange(2, 4),
			range2:   section.NewRange(6, 8),
			expected: false,
		},
		{
			range1:   section.NewRange(2, 3),
			range2:   section.NewRange(4, 5),
			expected: false,
		},
		{
			range1:   section.NewRange(5, 7),
			range2:   section.NewRange(7, 9),
			expected: true,
		},
		{
			range1:   section.NewRange(2, 8),
			range2:   section.NewRange(3, 7),
			expected: true,
		},
		{
			range1:   section.NewRange(6, 6),
			range2:   section.NewRange(4, 6),
			expected: true,
		},
		{
			range1:   section.NewRange(2, 6),
			range2:   section.NewRange(4, 8),
			expected: true,
		},
	}

	for i, s := range scenarios {
		overlapping1 := s.range1.Overlaps(s.range2)
		overlapping2 := s.range2.Overlaps(s.range1)

		if overlapping1 != s.expected || overlapping2 != s.expected {
			t.Logf("mismatched expectations at %v: expected %v, got %v, %v", i, s.expected, overlapping1, overlapping2)
			t.Fail()
		}
	}
}
