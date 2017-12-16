package main

import "fmt"
import "github.com/Wrenky/sudoKu/solve"

type Puzzle [][]uint

func main() {
	fmt.Println("Starting!")
	meanPuzzle := Puzzle{
		{0, 0, 0, 0, 0, 5, 0, 8, 0},
		{0, 0, 0, 6, 0, 1, 0, 4, 3},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 1, 0, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 6, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 0, 5},

		{5, 3, 0, 0, 0, 0, 0, 6, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Printf("Hard Puzzle starting state:\n%s\n", solve.Display(meanPuzzle))
	solved, _ := solve.SolvePuzzle(meanPuzzle)
	fmt.Printf("Hard Puzzle solved state:\n%s\n", solve.Display(solved))

}
