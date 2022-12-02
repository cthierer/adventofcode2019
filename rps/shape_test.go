package rps_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/rps"
)

func TestRockScoore(t *testing.T) {
	r := rps.Rock{}
	if r.Score() != 1 {
		t.Fail()
	}
}

func TestRockCompare(t *testing.T) {
	t.Parallel()
	r := rps.Rock{}

	t.Run("ties rock", func(t *testing.T) {
		if r.Compare(rps.Rock{}) != 0 {
			t.Fail()
		}
	})

	t.Run("loses to paper", func(t *testing.T) {
		if r.Compare(rps.Paper{}) != -1 {
			t.Fail()
		}
	})

	t.Run("beats scissors", func(t *testing.T) {
		if r.Compare(rps.Scissors{}) != 1 {
			t.Fail()
		}
	})
}

func TestPaperScore(t *testing.T) {
	p := rps.Paper{}
	if p.Score() != 2 {
		t.Fail()
	}
}

func TestPaperCompare(t *testing.T) {
	t.Parallel()
	p := rps.Paper{}

	t.Run("ties paper", func(t *testing.T) {
		if p.Compare(rps.Paper{}) != 0 {
			t.Fail()
		}
	})

	t.Run("loses to scissors", func(t *testing.T) {
		if p.Compare(rps.Scissors{}) != -1 {
			t.Fail()
		}
	})

	t.Run("beats rock", func(t *testing.T) {
		if p.Compare(rps.Rock{}) != 1 {
			t.Fail()
		}
	})
}

func TestScissorsScore(t *testing.T) {
	s := rps.Scissors{}
	if s.Score() != 3 {
		t.Fail()
	}
}

func TestScissorsCompare(t *testing.T) {
	t.Parallel()
	s := rps.Scissors{}

	t.Run("ties sciessors", func(t *testing.T) {
		if s.Compare(rps.Scissors{}) != 0 {
			t.Fail()
		}
	})

	t.Run("loses to rock", func(t *testing.T) {
		if s.Compare(rps.Rock{}) != -1 {
			t.Fail()
		}
	})

	t.Run("beats paper", func(t *testing.T) {
		if s.Compare(rps.Paper{}) != 1 {
			t.Fail()
		}
	})
}
