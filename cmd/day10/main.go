package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/cthierer/advent-of-code/device/cpu"
)

const (
	screenWidth = 40
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
		instr, err := parseLine(l)
		if err != nil {
			log.Fatalf("problem processing line: %v", err)
		}

		processor.Load(instr)
		valuesX[processor.Cycle()] = processor.X()
		cycle := float64(processor.Cycle())
		for processor.Next() {
			x := processor.X()
			valuesX[processor.Cycle()] = x
			cycle += 1
		}
	}

	totalSignalStrength := int64(0)
	for i := 20; i < processor.Cycle(); i += 40 {
		cycle := int64(i)
		x := valuesX[i]
		signalStrength := cycle * x
		totalSignalStrength += signalStrength
	}

	fmt.Printf("Total signal strength: %v\n", totalSignalStrength)

	fmt.Println("Screen output:")
	for i := 1; i < processor.Cycle(); i += 1 {
		row := int64(math.Floor(float64(i-1) / screenWidth))
		col := int64(i) - row*screenWidth - 1
		x := valuesX[i]

		if row > 0 && col == 0 {
			fmt.Println()
		}

		if x-1 <= col && col <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
