package test

import (
	"testing"

	playgroundManager "github.com/tictacgo/internal/manager/playgroundManager"
)

func TestPossibleGameResults(t *testing.T) {
	tictacPlayground := playgroundManager.NewPlayground()
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() and IsGameOver() returns false when queries in the middle of the game
	tictacPlayground.Input(0, 0, "X")
	tictacPlayground.Input(0, 1, "O")
	tictacPlayground.Input(0, 2, "X")
	tictacPlayground.Input(1, 0, "O")
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	if tictacPlayground.IsGameOver() != "" {
		t.Error("Game is over and we have a winner")
	}

	//Testing to see if IsFull() returns false and IsGameOver() returns true when X wins
	tictacPlayground.Input(1, 1, "X")
	tictacPlayground.Input(1, 2, "O")
	tictacPlayground.Input(2, 0, "X")

	if tictacPlayground.IsGameOver() != "X" {
		t.Error("O person has won the game when it should have been X")
	}
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() returns false and IsGameOver() returns true when X wins from a vertical column filled with X
	tictacPlayground.Input(0, 0, "O")
	tictacPlayground.Input(0, 1, "X")
	tictacPlayground.Input(0, 2, "O")
	tictacPlayground.Input(1, 0, "O")
	tictacPlayground.Input(1, 1, "X")
	tictacPlayground.Input(1, 2, "O")
	tictacPlayground.Input(2, 0, "X")
	tictacPlayground.Input(2, 1, "X")
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	if tictacPlayground.IsGameOver() != "X" {
		t.Error("O person has won the game when it should have been X")
	}
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() returns false and IsGameOver() returns true when O wins from a horizontal row filled with O
	tictacPlayground.Input(0, 0, "X")
	tictacPlayground.Input(0, 1, "X")
	tictacPlayground.Input(0, 2, "O")
	tictacPlayground.Input(1, 0, "O")
	tictacPlayground.Input(1, 1, "O")
	tictacPlayground.Input(1, 2, "O")
	tictacPlayground.Input(2, 0, "X")
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	if tictacPlayground.IsGameOver() != "O" {
		t.Error("X person has won the game when it should have been O")
	}
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() returns false and IsGameOver() returns true when X wins from a diagonal row filled with X
	tictacPlayground.Input(0, 0, "O")
	tictacPlayground.Input(0, 1, "O")
	tictacPlayground.Input(0, 2, "X")
	tictacPlayground.Input(1, 0, "0")
	tictacPlayground.Input(1, 1, "X")
	tictacPlayground.Input(1, 2, "O")
	tictacPlayground.Input(2, 0, "X")
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	if tictacPlayground.IsGameOver() != "X" {
		t.Error("O person has won the game when it should have been X")
	}
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() returns false and IsGameOver() returns true when O wins from a diagonal row filled with O from the other side
	tictacPlayground.Input(0, 0, "O")
	tictacPlayground.Input(0, 1, "X")
	tictacPlayground.Input(0, 2, "X")
	tictacPlayground.Input(1, 0, "X")
	tictacPlayground.Input(1, 1, "O")
	tictacPlayground.Input(2, 2, "O")
	if tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	if tictacPlayground.IsGameOver() != "O" {
		t.Error("X person has won the game when it should have been O")
	}
	tictacPlayground.ClearBoard()

	//Testing to see if IsFull() returns true and IsGameOver() returns false when no one wins
	tictacPlayground.Input(0, 0, "O")
	tictacPlayground.Input(0, 1, "X")
	tictacPlayground.Input(0, 2, "O")
	tictacPlayground.Input(1, 0, "X")
	tictacPlayground.Input(1, 1, "O")
	tictacPlayground.Input(1, 2, "X")
	tictacPlayground.Input(2, 0, "X")
	tictacPlayground.Input(2, 1, "O")
	tictacPlayground.Input(2, 2, "X")
	if tictacPlayground.IsGameOver() != "" {
		t.Error("Game is not over yet")
	}
	if !tictacPlayground.IsFull() {
		t.Error("Gameboard is actually full")
	}
	tictacPlayground.ClearBoard()

}

func TestInvalidGameInputs(t *testing.T) {

	tictacPlayground := playgroundManager.NewPlayground()
	tictacPlayground.ClearBoard()

	tictacPlayground.Input(0, 0, "X")
	tictacPlayground.Input(0, 1, "X")
	tictacPlayground.Input(0, 2, "O")
	err := tictacPlayground.Input(0, 2, "X")
	if !err {
		t.Error("The program is taking an invalid input: of repeated input")
	}

	err = tictacPlayground.Input(0, 1, "X")
	if !err {
		t.Error("The program is taking an invalid input: of repeated input")
	}

	err = tictacPlayground.Input(0, 3, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

	err = tictacPlayground.Input(0, -1, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

	err = tictacPlayground.Input(3, -1, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

	err = tictacPlayground.Input(-1, 3, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

	err = tictacPlayground.Input(-1, 1, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

	err = tictacPlayground.Input(3, 1, "X")
	if !err {
		t.Error("The program is taking an invalid input: of invalid input")
	}

}
