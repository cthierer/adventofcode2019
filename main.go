package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/inventory"
)

func selectMax(values []inventory.Tuple) inventory.Tuple {
	var max inventory.Tuple
	for _, v := range values {
		if v.Value() > max.Value() {
			max = v
		}
	}
	return max
}

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

	sums := inv.SumByIndex()
	max := selectMax(sums)

	fmt.Printf("Largest number of calories: %v (Elf #%v)\n", max.Value(), max.Index()+1)
}
