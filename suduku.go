package main

import (
	"fmt"
	"github.com/Wrenky/sudoKu/solve"
)

func main() {
	fmt.Println("Starting!")
	startingPuzzle := [][]uint{
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

	out := solve.Display(startingPuzzle)
	fmt.Println("Starting Puzzle:")
	fmt.Println("-------------------------------")
	fmt.Println(out)
	fmt.Println("-------------------------------")

	res, err := solve.SolvePuzzle(startingPuzzle)
	if err != nil {
		fmt.Println("Some unknown error during solving.")
		fmt.Println(err)
	} else {
		fmt.Println("Solved:")
		output := solve.Display(res)
		fmt.Println("-------------------------------")
		fmt.Println(output)
		fmt.Println("-------------------------------")
	}
}
