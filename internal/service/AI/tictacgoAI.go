package playgroundService

import (
	"fmt"
	"math/rand"
	"time"
)

type AI struct {
	player  string
	AILevel func(board [][]string) (int, int)
}

func (ai *AI) SetName(name string) {
	ai.player = name
}

func (ai *AI) AILevelEasy(board [][]string) (int, int) {
	return ai.GetSimpleMove(board)
}

func (ai *AI) AILevelMedium(board [][]string) (int, int) {
	return ai.GetRandomMove(board)
}

func (ai *AI) AILevelHard(board [][]string) (int, int) {
	//return ai.GetMinimaxMove(board)
	return ai.GetRandomMove(board)
}

func NewAI(player string) *AI {
	return &AI{player: player}
}

// Set AI level chooses
func (ai *AI) SetLevel(level int) {
	switch level {
	case 1:
		fmt.Println("Level 1 was selected")
		ai.AILevel = ai.AILevelEasy
	case 2:
		fmt.Println("Level 2 was selected")
		ai.AILevel = ai.AILevelMedium
	case 3:
		fmt.Println("Level 3 was selected")
		ai.AILevel = ai.AILevelHard
	default:
		fmt.Println("Level 1 was selected")
		ai.AILevel = ai.AILevelEasy
	}
}

func (ai *AI) GetSimpleMove(board [][]string) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return i, j
			}
		}
	}
	return -1, -1
}

func (ai *AI) GetRandomMove(board [][]string) (int, int) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	randomIndex2 := rand.Intn(3)
	if board[randomIndex][randomIndex2] == " " {
		return randomIndex, randomIndex2
	}
	return ai.GetRandomMove(board)
}

// func (ai *AI) GetMinimaxMove(board [][]string) (int, int) {

// }
