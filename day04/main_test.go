package main

import "testing"

func TestMatches(t *testing.T) {
	table := []struct {
		input    int
		expected bool
	}{
		{
			input:    111111,
			expected: true,
		},
		{
			input:    223450,
			expected: false,
		},
		{
			input:    123789,
			expected: false,
		},
	}

	for _, v := range table {
		actual := matches(v.input)
		if actual != v.expected {
			t.Fail()
		}
	}
}

func TestMatches2(t *testing.T) {
	table := []struct {
		input    int
		expected bool
	}{
		{
			input:    112233,
			expected: true,
		},
		{
			input:    123444,
			expected: false,
		},
		{
			input:    111122,
			expected: true,
		},
	}

	for _, v := range table {
		actual := matches2(v.input)
		if actual != v.expected {
			t.Fail()
		}
	}
}
