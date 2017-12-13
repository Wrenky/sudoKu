package main

import (
	"fmt"
	"github.com/Wrenky/sudoKu/solve"
	"reflect"
)

type Puzzle [][]uint

func main() {
	fmt.Println("Starting!")
	easyPuzzle := Puzzle{
		{0, 0, 0, 6, 0, 4, 7, 0, 0},
		{7, 0, 6, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 8, 0},
		{0, 7, 0, 0, 2, 0, 0, 9, 3},
		{8, 0, 0, 0, 0, 0, 0, 0, 5},
		{4, 3, 0, 0, 1, 0, 0, 7, 0},
		{0, 5, 0, 2, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 2, 0, 8},
		{0, 0, 2, 3, 0, 1, 0, 0, 0},
	}

	hardPuzzle := Puzzle{
		{0, 0, 0, 0, 1, 0, 9, 0, 0},
		{9, 0, 0, 7, 0, 2, 3, 0, 0},
		{0, 5, 0, 8, 0, 9, 0, 0, 0},

		{0, 0, 5, 0, 3, 8, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 9},
		{0, 9, 0, 1, 2, 0, 6, 0, 0},

		{0, 0, 0, 2, 0, 7, 0, 3, 0},
		{0, 0, 6, 3, 0, 4, 0, 0, 2},
		{0, 0, 8, 0, 6, 0, 0, 0, 0},
	}

	carlPuzzle := Puzzle{
		{7, 3, 4, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 9, 2, 0},
		{0, 0, 0, 0, 8, 5, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 1, 3, 8},
		{0, 0, 0, 4, 7, 2, 0, 0, 0},
		{5, 9, 6, 0, 0, 0, 0, 0, 0},
		{3, 0, 0, 9, 6, 0, 0, 0, 0},
		{0, 1, 2, 8, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 5, 4},
	}

	puzzleList := []Puzzle{easyPuzzle, hardPuzzle, carlPuzzle}

	fmt.Println("----------------------------------------------------------")
	for index, puzzle := range puzzleList {
		out := solve.Display(puzzle)
		fmt.Printf("Puzzle #%d initial state:\n%s\n", index, out)
		res, err := solve.SolvePuzzle(puzzle)
		if err != nil {
			fmt.Println("Some unknown error during solving.")
			fmt.Println(err)
		} else {
			fmt.Println("Solved:")
			output := solve.Display(res)
			fmt.Println(output)
		}
		fmt.Println("----------------------------------------------------------")
		hintedPuzzle, row, col, err := solve.Hint(puzzle)
		fmt.Println("Type of err: ", reflect.TypeOf(err))
		if pe, ok := err.(*solve.PuzzleError); ok {
			fmt.Printf("Puzzled! At (%d,%d): --%s--", pe.Row, pe.Col, pe)
		}
		fmt.Printf("Getting a hint, post hint puzzle(%d,%d):\n", row, col)
		output := solve.Display(hintedPuzzle)
		fmt.Println(output)
	}
}
