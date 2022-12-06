package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cthierer/advent-of-code/cargo"
	"github.com/cthierer/advent-of-code/crane"
)

func parseInput(input string) (crates string, commands string, err error) {
	separatorIdx := 0
	lines := strings.Split(input, "\n")
	for separatorIdx = 0; separatorIdx < len(lines); separatorIdx++ {
		if lines[separatorIdx] == "" {
			break
		}
	}

	if separatorIdx == len(lines) {
		err = errors.New("invalid input format")
		return
	}

	crates = strings.Join(lines[0:separatorIdx], "\n")
	commands = strings.Join(lines[separatorIdx+1:], "\n")
	return
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	crates, commands, err := parseInput(string(input))
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	collection, err := cargo.ParseCollection(crates)
	if err != nil {
		log.Fatalf("failed to build collection: %v", err)
	}

	cmds, err := crane.ParseCommands(commands)
	if err != nil {
		log.Fatalf("failed to build commands: %v", err)
	}

	for _, cmd := range cmds {
		collection.Transfer(cmd)
	}

	stacks := collection.Values()
	topValues := make([]string, len(stacks))
	for i, s := range stacks {
		topValues[i] = s.Peek().String()
	}

	fmt.Printf("Top of each stack: %v\n", strings.Join(topValues, ""))
}
