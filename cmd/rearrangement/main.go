package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cthierer/advent-of-code/rucksack"
)

func parseInput(input string) []*rucksack.Rucksack {
	lines := strings.Split(input, "\n")
	rucksacks := make([]*rucksack.Rucksack, len(lines))
	for i, l := range lines {
		rucksacks[i] = rucksack.ParseRucksack(l)
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

	groups := make([]*rucksack.Group, len(rucksacks)/3)
	j := 0
	for i := 0; i < len(rucksacks); i += 3 {
		g := rucksack.Group{}
		g.Add(rucksacks[i])
		g.Add(rucksacks[i+1])
		g.Add(rucksacks[i+2])
		groups[j] = &g
		j += 1
	}

	uniqueGroupItems := rucksack.NewItemCollection()
	for _, g := range groups {
		uniqueGroupItems.Join(g.OverlappingItemTypes())
	}

	fmt.Printf("Sum of overlapping group item type priorities: %v\n", uniqueGroupItems.Sum())
}
