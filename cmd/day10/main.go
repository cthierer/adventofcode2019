package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cthierer/advent-of-code/device/cpu"
)

func parseLine(l string) (cpu.Instruction, error) {
	parts := strings.Split(l, " ")
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "addx":
		value, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return nil, err
		}
		return &cpu.AddX{Value: value}, nil
	case "noop":
		return &cpu.Noop{}, nil
	default:
		return nil, fmt.Errorf("unrexognized command: %v", cmd)
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	processor := &cpu.Processor{}
	lines := strings.Split(string(input), "\n")
	valuesX := make(map[int]int64)
	for _, l := range lines {
		i, err := parseLine(l)
		if err != nil {
			log.Fatalf("problem processing line: %v", err)
		}

		log.Printf("loading: %v", i)
		processor.Load(i)
		for processor.Next() {
			valuesX[processor.Cycle()] = processor.X()
			log.Printf("status: %v", processor)
		}
	}

	totalSignalStrength := int64(0)
	for i := 19; i < len(valuesX); i += 40 {
		cycle := int64(i + 1)
		x := valuesX[i]
		signalStrength := cycle * x
		totalSignalStrength += signalStrength
		log.Printf("cycle=%v, value=%v, strength=%v", cycle, valuesX[i], signalStrength)
	}

	fmt.Printf("Total signal strength: %v\n", totalSignalStrength)
}
