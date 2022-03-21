package playgroundManager

import (
	playgroundAI "github.com/tictacgo/internal/service/AI"
	playground "github.com/tictacgo/internal/service/playground"
)

var tictacGoAI = playgroundAI.NewAI("O")

//Set the AI name
func SetAIName(name string) {
	tictacGoAI.SetName(name)
}

func GetAIMove(board *playground.Board) (int, int) {
	r, c := tictacGoAI.AILevel(board.GetBoard())
	return r, c
}

func SetAILevel(level int) {
	tictacGoAI.SetLevel(level)
}
