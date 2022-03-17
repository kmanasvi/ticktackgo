package playgroundManager

import (
	playgroundAI "github.com/tictacgo/internal/service/AI"
	playground "github.com/tictacgo/internal/service/playground"
)

var tictacGoAI = playgroundAI.NewAI("Computer")

func GetAIMove(board *playground.Board) (int, int) {
	return tictacGoAI.GetRandomMove(board)
}
