package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func traverseUp(forest [][]int64, x, y int, threshold int64, count int) int {
	if y < 0 {
		return count
	}

	count += 1

	if forest[y][x] >= threshold {
		return count
	}

	return traverseUp(forest, x, y-1, threshold, count)
}

func traverseDown(forest [][]int64, x, y int, threshold int64, count int) int {
	if y >= len(forest) {
		return count
	}

	count += 1

	if forest[y][x] >= threshold {
		return count
	}

	return traverseDown(forest, x, y+1, threshold, count)
}

func traverseLeft(forest [][]int64, x, y int, threshold int64, count int) int {
	if x < 0 {
		return count
	}

	count += 1

	if forest[y][x] >= threshold {
		return count
	}

	return traverseLeft(forest, x-1, y, threshold, count)
}

func traverseRight(forest [][]int64, x, y int, threshold int64, count int) int {
	if x >= len(forest[y]) {
		return count
	}

	count += 1

	if forest[y][x] >= threshold {
		return count
	}

	return traverseRight(forest, x+1, y, threshold, count)
}

func scenicScore(forest [][]int64, x, y int) int {
	height := forest[y][x]
	scoreUp := traverseUp(forest, x, y-1, height, 0)
	scoreDown := traverseDown(forest, x, y+1, height, 0)
	scoreLeft := traverseLeft(forest, x-1, y, height, 0)
	scoreRight := traverseRight(forest, x+1, y, height, 0)
	return scoreUp * scoreDown * scoreLeft * scoreRight
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
	maxScore := 0

	for i := 1; i < height-1; i += 1 {
		for j := 1; j < width-1; j += 1 {
			score := scenicScore(forest, j, i)
			log.Printf("tree at col=%v, row=%v has score %v", j+1, i+1, score)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Printf("Maximum scenic score: %v\n", maxScore)
}
