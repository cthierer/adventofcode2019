package intcode

import (
	"testing"
)

func TestRun(t *testing.T) {
	table := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 0, 0, 0, 99},
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			input:    []int{2, 3, 0, 3, 99},
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			input:    []int{2, 4, 4, 5, 99, 0},
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			input:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, r := range table {
		actual, err := Run(r.input)
		if err != nil {
			t.Fatal()
		}
		for i, v := range r.expected {
			if actual[i] != v {
				t.Fail()
				break
			}
		}
	}
}
