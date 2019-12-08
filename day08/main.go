package main

import (
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func getLayer(pixels []int, layerSize, idx int) []int {
	startAt := idx * layerSize
	endAt := startAt + layerSize
	return pixels[startAt:endAt]
}

func findLastLayer(pixels []int, layerSize int, cb func([]int) bool) int {
	numLayers := len(pixels) / layerSize
	lastLayerIdx := 0

	for i := 0; i < numLayers; i++ {
		layer := getLayer(pixels, layerSize, i)
		if cb(layer) {
			lastLayerIdx = i
		}
	}

	return lastLayerIdx
}

func countOccurrencesOf(val int) func([]int) int {
	return func(in []int) int {
		count := 0
		for _, v := range in {
			if v == val {
				count++
			}
		}
		return count
	}
}

func layerWithFewest(count func([]int) int) func([]int) bool {
	fewest := -1
	return func(pixels []int) bool {
		count := count(pixels)
		if fewest < 0 || count < fewest {
			fewest = count
			return true
		}
		return false
	}
}

func solvePuzzle1(pixels []int, width, height int) (int, int) {
	layerSize := width * height

	countZeros := countOccurrencesOf(0)
	findLayer := layerWithFewest(countZeros)
	layerIdx := findLastLayer(pixels, layerSize, findLayer)

	countOnes := countOccurrencesOf(1)
	countTwos := countOccurrencesOf(2)
	layer := getLayer(pixels, layerSize, layerIdx)
	solution := countOnes(layer) * countTwos(layer)

	return layerIdx, solution
}

func toArray(input string) []int {
	arr := make([]int, utf8.RuneCountInString(input))
	for i, r := range input {
		arr[i] = int(r - '0')
	}
	return arr
}

func main() {
	width := 25
	height := 6
	contents, err := ioutil.ReadFile("./day08/input")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	input := toArray(string(contents))
	layer, solution := solvePuzzle1(input, width, height)
	fmt.Printf("layer #%d, solution = %d\n", layer+1, solution)
}
