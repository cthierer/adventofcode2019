# Rock, paper, scissors

Plays a game of rock, paper, scissors using the moves described in 
[the strategy guide](https://adventofcode.com/2022/day/2):

> The first column is what your opponent is going to play: A for Rock, B for Paper, and C for 
> Scissors. The second column -- X means you need to lose, Y means you need to end the round in a 
> draw, and Z means you need to win.

The file is read in from `stdin` up until `EOF`.

Score is computed using the score guide provided:

> Your total score is the sum of your scores for each round. The score for a single round is the 
> score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for 
> the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

## Sample usage

```bash
go run cmd/rockpaperscissors/main.go < input-day02.txt
```
