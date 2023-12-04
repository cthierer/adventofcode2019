package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/trebuchet"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	document := trebuchet.ParseDocument(string(input))

	result, err := trebuchet.SumCalibrationValues(document)
	if err != nil {
		log.Fatalf("failed to sum calibration values: %v", err)
	}

	fmt.Printf("Sum of calibration values: %v\n", result)
}
