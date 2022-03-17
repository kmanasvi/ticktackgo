package playgroundService

import (
	"time"
)

type Board struct {
	uuid  string
	board [][]string
}

func NewBoard() *Board {
	board := make([][]string, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]string, 3)
	}
	return &Board{time.Now().String(), board}
}

func (b *Board) GetBoard() [][]string {
	return b.board
}

func (b *Board) IsGameOver() string {

	for i := 0; i < len(b.board); i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] != " " {
			return b.board[i][0]
		}

		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] != " " {
			return b.board[0][i]
		}
	}

	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[0][2] != " " {
		return b.board[0][2]
	}

	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[0][0] != " " {
		return b.board[0][0]
	}

	return ""
}

func (b *Board) Input(row int, col int, character string) bool {
	if row >= 0 && row < 3 && col >= 0 && col < 3 && b.board[row][col] == " " {
		b.board[row][col] = character
		return false
	} else {
		return true
	}
}

func (b *Board) ClearBoard() {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board); j++ {
			b.board[i][j] = " "
		}
	}
}

func (b *Board) IsFull() bool {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
