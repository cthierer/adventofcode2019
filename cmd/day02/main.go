package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cthierer/advent-of-code/cubegame"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	games, err := cubegame.ParseGameSet(string(input))
	if err != nil {
		log.Fatalf("failed to parse game report: %v", err)
	}

	query := cubegame.NewCubeSet(12, 13, 14)
	possibleGames := cubegame.QueryPossibleGames(games, query)

	fmt.Printf("Possible games (%v total):\n%v\n", len(possibleGames.Games), possibleGames)

	idSum := 0
	for _, game := range possibleGames.Games {
		idSum += game.ID
	}

	fmt.Printf("Sum of possible game IDs: %v\n", idSum)

	powerSum := 0
	for _, game := range games.Games {
		minCubes := game.MinCubes()
		power := minCubes.Power()
		powerSum += power
	}

	fmt.Printf("Sum of the power of the minimum sets: %v\n", powerSum)
}
