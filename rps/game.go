package rps

type Game struct {
	scorePlayer1 int
	scorePlayer2 int
}

type Round interface {
	Player1() Action
	Player2() Action
}

const (
	POINTS_TIE  = 3
	POINTS_WON  = 6
	POINTS_LOST = 0
)

func (g *Game) Play(r Round) {
	scorePlayer1 := r.Player1().Score()
	scorePlayer2 := r.Player2().Score()

	player1Result := r.Player1().Compare(r.Player2())
	switch player1Result {
	case -1:
		scorePlayer1 += POINTS_LOST
		scorePlayer2 += POINTS_WON
	case 1:
		scorePlayer1 += POINTS_WON
		scorePlayer2 += POINTS_LOST
	default:
		scorePlayer1 += POINTS_TIE
		scorePlayer2 += POINTS_TIE
	}

	g.scorePlayer1 += scorePlayer1
	g.scorePlayer2 += scorePlayer2
}

func (g *Game) Scores() (player1, player2 int) {
	return g.scorePlayer1, g.scorePlayer2
}
