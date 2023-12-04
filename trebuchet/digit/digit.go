package digit

import (
	"errors"
	"strings"
)

var ErrNotFound = errors.New("no digit found")

var literals = map[string]int{
	"0":     0,
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

func FindFirstOccurrence(search string) (int, error) {
	var value int
	idx := -1

	for literal, digit := range literals {
		foundIdx := strings.Index(search, literal)
		if foundIdx < 0 {
			continue
		}
		if idx < 0 || foundIdx < idx {
			value = digit
			idx = foundIdx
		}
	}

	if idx < 0 {
		return 0, ErrNotFound
	}

	return value, nil
}

func FindLastOccurrence(search string) (int, error) {
	var value int
	idx := -1

	for literal, digit := range literals {
		foundIdx := strings.LastIndex(search, literal)
		if foundIdx < 0 {
			continue
		}
		if foundIdx > idx {
			value = digit
			idx = foundIdx
		}
	}

	if idx < 0 {
		return 0, ErrNotFound
	}

	return value, nil
}
