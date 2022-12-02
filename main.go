package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/inventory"
)

// usage: go run main.go < input.txt
func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	inv, err := inventory.ParseInventory(string(input))
	if err != nil {
		log.Fatalf("failed to parse inentory: %v", err)
	}

	topThree := inv.SumByIndex().Sort().Slice(0, 3)

	fmt.Printf("Top %v calorie counts:\n", len(topThree.LineItems()))
	for i, t := range topThree.LineItems() {
		fmt.Printf("%v.)\tElf #%5v\t%v calories\n", i+1, t.Index()+1, t.Value())
	}

	total := topThree.Total()
	fmt.Printf("Total calories: %v\n", total)
}
