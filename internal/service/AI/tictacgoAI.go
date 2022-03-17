package playgroundService

import (
	"math/rand"
	"time"

	playground "github.com/tictacgo/internal/service/playground"
)

type AI struct {
	player string
}

func NewAI(player string) *AI {
	return &AI{player}
}

func (ai *AI) GetSimpleMove(board *playground.Board) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetBoard()[i][j] == " " {
				return i, j
			}
		}
	}
	return -1, -1
}

func (ai *AI) GetRandomMove(board *playground.Board) (int, int) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	randomIndex2 := rand.Intn(3)
	if board.GetBoard()[randomIndex][randomIndex2] == " " {
		return randomIndex, randomIndex2
	}
	return ai.GetRandomMove(board)
}

//TODO - Implement minimax algorithm to determine the best move for the AI
// func (ai *AI) GetMiniMaxMove(board *playground.Board) (int, int) {

// }
