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

	p := signal.FindStartOfPacket(string(input))
	if p == -1 {
		log.Fatalf("failed to find start of packet")
	}

	fmt.Printf("Packet starts after %v characters\n", p)

	m := signal.FindStartOfMessage(string(input))
	if m == -1 {
		log.Fatalf("failed to find start of message")
	}

	fmt.Printf("Message starts after %v characters\n", m)
}
