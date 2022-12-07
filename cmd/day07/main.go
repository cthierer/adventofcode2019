package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/device/shell"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	s, err := shell.Parse(string(input))
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	sum := int64(0)
	for _, dir := range s.FS.Dirs() {
		if dir.TotalSize() <= 100000 {
			sum += dir.TotalSize()
		}
	}

	fmt.Printf("Total size of directories with at most 100000: %v\n", sum)
}
