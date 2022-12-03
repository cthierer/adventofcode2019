package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cthierer/advent-of-code/rucksack"
)

func addContents(contents string, compartment *rucksack.ItemCollection) {
	for _, r := range contents {
		compartment.Add(rucksack.ParseItemType(r))
	}
}

func parseLine(line string) *rucksack.Rucksack {
	r := rucksack.NewRucksack()
	halfway := len(line) / 2
	compartment1 := line[0:halfway]
	compartment2 := line[halfway:]

	addContents(compartment1, r.Compartment1)
	addContents(compartment2, r.Compartment2)

	return r
}

func parseInput(input string) []*rucksack.Rucksack {
	lines := strings.Split(input, "\n")
	rucksacks := make([]*rucksack.Rucksack, len(lines))
	for i, l := range lines {
		rucksacks[i] = parseLine(l)
	}
	return rucksacks
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	rucksacks := parseInput(string(input))
	uniqueItems := rucksack.NewItemCollection()

	for _, r := range rucksacks {
		uniqueItems.Join(r.OverlappingItemTypes())
	}

	fmt.Printf("Sum of overlapping item type priorities: %v\n", uniqueItems.Sum())
}
