package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/rps"
	"github.com/cthierer/advent-of-code/strategyguide"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	guide, err := strategyguide.ParseStrategyGuide(string(input))
	if err != nil {
		log.Fatalf("failed to parse guide: %v", err)
	}

	game := rps.Game{}
	for guide.Next() {
		turn := guide.Get()
		if turn == strategyguide.NoEntry {
			log.Fatal("invalid turn")
		}
		game.Play(turn)
	}

	player1, player2 := game.Scores()
	fmt.Println("Final scores")
	fmt.Printf("Player 1\t%v\nPlayer 2\t%v\n", player1, player2)
}
