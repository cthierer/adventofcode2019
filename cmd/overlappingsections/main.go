package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cthierer/advent-of-code/section"
)

func parseInput(input string) (pairs []section.Pair, err error) {
	lines := strings.Split(input, "\n")
	pairs = make([]section.Pair, len(lines))
	for i, l := range lines {
		pairs[i], err = section.ParsePair(l)
		if err != nil {
			return nil, err
		}
	}
	return pairs, nil
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	pairs, err := parseInput(string(input))
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	numFullyOverlapping := 0
	numPartiallyOverlapping := 0
	for _, p := range pairs {
		if p.FullyOverlapping() {
			numFullyOverlapping += 1
		}
		if p.ParitallyOverlapping() {
			numPartiallyOverlapping += 1
		}
	}

	fmt.Printf("Num. of pairs that have overlapping ranges: %v\n", numFullyOverlapping)
	fmt.Printf("Num. of paris that have paritally overlapping ranges: %v\n", numPartiallyOverlapping)
}
