package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/signal"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	s := signal.FindStartOfPacket(string(input))
	if s == -1 {
		log.Fatalf("failed to find start of packet")
	}

	fmt.Printf("Packet starts after %v characters\n", s)
}
