package main

import "fmt"
import "github.com/Wrenky/sudoKu/solve"

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

	fmt.Printf("Easy Puzzle starting state:\n%s\n", solve.Display(easyPuzzle))
	solved, _ := solve.SolvePuzzle(easyPuzzle)
	fmt.Printf("Easy Puzzle solved state:\n%s\n", solve.Display(solved))

}
