package playgroundManager

//TODO - Implement 2 player mode
// TODO - Implement difficulty level to play against AI
// TODO - Easy being simple AI, Medium being random AI, Hard being minimax AI
// TODO - Implement a scoring system to determine who won the game
// TODO - Implement a leaderboard to determine the top players
// TODO - Implement a system which allows the user to save their game and load it later
// TODO - Allow the players to pick either X or O when playing with the AI

import (
	"fmt"

	tictacgoAIManager "github.com/tictacgo/internal/manager/AI"
	playground "github.com/tictacgo/internal/service/playground"
)

var tictacPlayground = playground.NewBoard()

func NewPlayground() *playground.Board {
	return tictacPlayground
}

func InitilizeBoard() {
	fmt.Println("Press s to start game, q to quit")
	var input string
	fmt.Scanln(&input)
	if input == "s" {
		tictacPlayground = playground.NewBoard()
		StartNewGame()
	} else if input == "q" {
		fmt.Println("Exiting")
	} else {
		fmt.Println("Invalid input")
		InitilizeBoard()
	}
}

func PickCharacter() string {
	fmt.Println("Player 1, Press x to play as X, o to play as O")
	var input string
	fmt.Scanln(&input)
	if input == "x" {
		return "X"
	}
	if input == "o" {
		return "O"
	}
	fmt.Println("Invalid input")
	return PickCharacter()
}

func SelectGameModes() string {
	fmt.Println("Press 1 to play single player, 2 to play two player")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		return "1"
	}
	if input == "2" {
		return "2"
	}
	fmt.Println("Invalid input")
	return SelectGameModes()
}

func getAILevelINput() {
	fmt.Println("Press 1 for Easy AI, 2 for Medium AI, 3 for Hard AI")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		tictacgoAIManager.SetAILevel(1)
		return
	}
	if input == "2" {
		tictacgoAIManager.SetAILevel(2)
		return
	}
	if input == "3" {
		tictacgoAIManager.SetAILevel(3)
		return
	}
	fmt.Println("Invalid input")
	getAILevelINput()
}

func getMove(character string) {
	err := true
	for {
		row, col := getPlayersMove()
		err = tictacPlayground.Input(row, col, character)
		if err {
			fmt.Println("Invalid location inputted")
			continue
		}
		break
	}
}

func getPlayersMove() (int, int) {
	fmt.Println("Enter the location of the character in the format r c where 0<=r<=2 and 0<=c<=2\n")
	var row, col int
	fmt.Scanf("%d %d", &row, &col)
	return row, col
}

func DisplayBoard(board [][]string) {
	for i := 0; i < 3; i++ {
		fmt.Print(board[i])
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func switchUser(character string) string {
	if character == "X" {
		return "O"
	}
	return "X"
}

func checkResult() string {
	result := tictacPlayground.IsGameOver()
	if result == "X" || result == "O" {
		fmt.Printf("Player %s wins!\n", result)
		return result
	}
	if tictacPlayground.IsFull() {
		fmt.Printf("Game is a draw!\n")
		return "draw"
	}
	return ""
}

func StartNewGame() {
	tictacPlayground.ClearBoard()
	gameMode := SelectGameModes()

	if gameMode == "1" {
		getAILevelINput()
	}

	character := PickCharacter()
	fmt.Println("Player 1, you are playing as ", character)
	tictacgoAIManager.SetAIName(switchUser(character))
	var result string

	for {

		if character == "X" {
			getMove(character)
			DisplayBoard(tictacPlayground.GetBoard())

			result = checkResult()
			if result != "" {
				break
			}

			character = switchUser(character)
			if gameMode == "1" {
				row, col := tictacgoAIManager.GetAIMove(tictacPlayground)
				tictacPlayground.Input(row, col, character)
			} else if gameMode == "2" {
				getMove(character)
			}
			DisplayBoard(tictacPlayground.GetBoard())
			result = checkResult()
			if result != "" {
				break
			}
			character = switchUser(character)

		} else {
			if gameMode == "1" {
				row, col := tictacgoAIManager.GetAIMove(tictacPlayground)
				tictacPlayground.Input(row, col, character)
			} else if gameMode == "2" {
				getMove(character)
			}
			DisplayBoard(tictacPlayground.GetBoard())
			result = checkResult()
			if result != "" {
				break
			}
			character = switchUser(character)
			getMove(character)
			DisplayBoard(tictacPlayground.GetBoard())
			result = checkResult()
			if result != "" {
				break
			}
			character = switchUser(character)
		}
	}
	InitilizeBoard()
}
