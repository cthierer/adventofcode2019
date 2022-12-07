package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/device/filesystem"
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

	freeSpace := 70000000 - s.FS.Root().TotalSize()
	needToFree := 30000000 - freeSpace

	var smallestDirInThreshold *filesystem.Dir
	for _, dir := range s.FS.Dirs() {
		size := dir.TotalSize()
		if size >= needToFree {
			if smallestDirInThreshold == nil || size < smallestDirInThreshold.TotalSize() {
				smallestDirInThreshold = dir
			}
		}
	}

	fmt.Printf("Smallest directory that can free up enough space (need %v): %s (%v)\n", needToFree, smallestDirInThreshold.Name, smallestDirInThreshold.TotalSize())
}
