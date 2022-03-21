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

//Returns the playground service
func NewPlayground() *playground.Board {
	return tictacPlayground
}

//Takes the user input to start a new game
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

func checkResult() string {
	//There may be a case where the game is over and the board is full as well so we want to check for that
	result := tictacPlayground.IsGameOver()
	if result == "X" || result == "O" {
		fmt.Println("Game is over")
		fmt.Println(result + " is the winner")
		return result
	}
	if tictacPlayground.IsFull() {
		fmt.Println("Game is a draw")
		return "draw"
	}
	return ""
}

//Clears the board and starts a new game. Initializes the board when a game is finished
func StartNewGame() {
	tictacPlayground.ClearBoard()
	character := "X"
	err := true
	var result string
	for {
		character = "X"
		for {
			row, col := getPlayersMove()
			err = tictacPlayground.Input(row, col, character)
			if err {
				fmt.Println("Invalid location inputted")
				continue
			}
			break
		}

		DisplayBoard(tictacPlayground.GetBoard())
		result = checkResult()
		if result != "" {
			break
		}
		character = "O"
		row, col := tictacgoAIManager.GetAIMove(tictacPlayground)
		tictacPlayground.Input(row, col, character)
		DisplayBoard(tictacPlayground.GetBoard())
		result = checkResult()
		if result != "" {
			break
		}
	}
	InitilizeBoard()
}

//Takes the user input for the row and column
func getPlayersMove() (int, int) {
	fmt.Println("Enter the location of the character in the format r c where 0<=r<=2 and 0<=c<=2\n")
	var row, col int
	fmt.Scanf("%d %d", &row, &col)
	return row, col
}

//Displays the board
func DisplayBoard(board [][]string) {
	for i := 0; i < 3; i++ {
		fmt.Print(board[i])
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
