package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func traverseUp(forest [][]int64, x, y int, tallest int64) int64 {
	if y < 0 {
		return tallest
	}

	if forest[y][x] >= tallest {
		tallest = forest[y][x]
	}

	return traverseUp(forest, x, y-1, tallest)
}

func traverseDown(forest [][]int64, x, y int, tallest int64) int64 {
	if y >= len(forest) {
		return tallest
	}

	if forest[y][x] >= tallest {
		tallest = forest[y][x]
	}

	return traverseDown(forest, x, y+1, tallest)
}

func traverseLeft(forest [][]int64, x, y int, tallest int64) int64 {
	if x < 0 {
		return tallest
	}

	if forest[y][x] >= tallest {
		tallest = forest[y][x]
	}

	return traverseLeft(forest, x-1, y, tallest)
}

func traverseRight(forest [][]int64, x, y int, tallest int64) int64 {
	if x >= len(forest[y]) {
		return tallest
	}

	if forest[y][x] >= tallest {
		tallest = forest[y][x]
	}

	return traverseRight(forest, x+1, y, tallest)
}

func isVisible(forest [][]int64, x, y int) bool {
	height := forest[y][x]
	if height > traverseUp(forest, x, y-1, int64(0)) {
		return true
	}

	if height > traverseDown(forest, x, y+1, int64(0)) {
		return true
	}

	if height > traverseLeft(forest, x-1, y, int64(0)) {
		return true
	}

	if height > traverseRight(forest, x+1, y, int64(0)) {
		return true
	}

	return false
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	lines := strings.Split(string(input), "\n")
	forest := make([][]int64, len(lines))

	for i, line := range lines {
		row := make([]int64, len(line))

		for j, v := range line {
			height, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				log.Fatalf("failed to parse height: %v", err)
			}

			row[j] = height
		}

		forest[i] = row
	}

	height := len(lines)
	width := len(lines[0])
	countVisible := (height+width)*2 - 4 // outer edge is always visible
	log.Printf("outer trees are visible: %v", countVisible)

	for i := 1; i < height-1; i += 1 {
		for j := 1; j < width-1; j += 1 {
			if isVisible(forest, j, i) {
				log.Printf("tree at col=%v, row=%v is visible", j+1, i+1)
				countVisible += 1
			}
		}
	}

	fmt.Printf("Number of trees visible from the edge: %v\n", countVisible)
}
