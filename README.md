<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [sudoKu](#sudoku)
  - [Getting Started](#getting-started)
- [Exposed Functions and Types](#exposed-functions-and-types)
  - [SolvePuzzle](#solvepuzzle)
  - [Hint](#hint)
  - [Display](#display)
  - [PuzzleError](#puzzleerror)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# sudoKu
A library to help solve sudoKu puzzles, written in Go. Implements [Peter Norvig's ideas about constraint propagation and search applied to sudoku.](http://norvig.com/sudoku.html) Its a great read, I highly recommend it!
Created for CSE682, and intended to work with Ryan Milbourne's API [syr-dudoku-backend](https://github.com/ryanbmilbourne/syr-sudoku-backend) and Carl Poole's android app.


## Getting Started
Simply just import this module, then you are ready to start using its methods!


# Exposed Functions and Types

## SolvePuzzle
 ```SolvePuzzle(puzzle [][]uint) (solvedPuzzle [][]uint,err error)```
SolvePuzzle is your main call into sudoKu! it takes a 2d slice of uints and returns a solved 2d slice of uints, or an error. 
For example, this block:
```
package main
import (
	"fmt"
	"github.com/Wrenky/sudoKu/solve"
)
type Puzzle [][]uint
func main() {
	puzzle := Puzzle{
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

	fmt.Printf("Starting puzzle:\n%s\n", solve.Display(puzzle))
	res, err := solve.SolvePuzzle(puzzle)
	if err != nil {
		fmt.Printf("Failed to solve the puzzle: %s", err)
	} else {
		fmt.Printf("Solved puzzle\n%s\n", solve.Display(res))
	}

	return
}

```
Produces this output:
```
Starting puzzle:
0  0  0 |  6  0  4 |  7  0  0
7  0  6 |  0  0  0 |  0  0  9
0  0  0 |  0  0  5 |  0  8  0
---------------------------------
0  7  0 |  0  2  0 |  0  9  3
8  0  0 |  0  0  0 |  0  0  5
4  3  0 |  0  1  0 |  0  7  0
---------------------------------
0  5  0 |  2  0  0 |  0  0  0
3  0  0 |  0  0  0 |  2  0  8
0  0  2 |  3  0  1 |  0  0  0

Solved puzzle
5  8  3 |  6  9  4 |  7  2  1
7  1  6 |  8  3  2 |  5  4  9
2  9  4 |  1  7  5 |  3  8  6
---------------------------------
6  7  1 |  5  2  8 |  4  9  3
8  2  9 |  7  4  3 |  1  6  5
4  3  5 |  9  1  6 |  8  7  2
---------------------------------
1  5  8 |  2  6  7 |  9  3  4
3  6  7 |  4  5  9 |  2  1  8
9  4  2 |  3  8  1 |  6  5  7

```


## Hint
 ```Hint(puzzle [][]uint) (updatedPuzzle [][]uint, row uint, col uint, err error)```
Hint takes a puzzle, and returns and updated puzzle and the row/col values + an error. This error can be a PuzzleErorr!
```
    puzzle := Puzzle{
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

    fmt.Printf("Pre-hint puzzle:\n%s\n", solve.Display(puzzle))
    hintedPuzzle, row, col, err := solve.Hint(puzzle)
    if pe, ok := err.(*solve.PuzzleError); ok {
        fmt.Printf("Puzzled! At (%d,%d): %s", pe.Row, pe.Col, pe)
    } else {
        fmt.Printf("Getting a hint, post hint puzzle(%d,%d):\n%s\n", row, col, solve.Display(hintedPuzzle))
    }

    return

```
In the above example,  we simpily want to get a hint. This returns:

```
  0  0  0 |  6  0  4 |  7  0  0
  7  0  6 |  0  0  0 |  0  0  9
  0  0  0 |  0  0  5 |  0  8  0
---------------------------------
  0  7  0 |  0  2  0 |  0  9  3
  8  0  0 |  0  0  0 |  0  0  5
  4  3  0 |  0  1  0 |  0  7  0
---------------------------------
  0  5  0 |  2  0  0 |  0  0  0
  3  0  0 |  0  0  0 |  2  0  8
  0  0  2 |  3  0  1 |  0  0  0

Getting a hint, post hint puzzle(8,0):
  0  0  0 |  6  0  4 |  7  0  0
  7  0  6 |  0  0  0 |  0  0  9
  0  0  0 |  0  0  5 |  0  8  0
---------------------------------
  0  7  0 |  0  2  0 |  0  9  3
  8  0  0 |  0  0  0 |  0  0  5
  4  3  0 |  0  1  0 |  0  7  0
---------------------------------
  0  5  0 |  2  0  0 |  0  0  0
  3  0  0 |  0  0  0 |  2  0  8
  9  0  2 |  3  0  1 |  0  0  0
```
indexes are 0-indexed, so check the bottom-left corner.

## Display
 ```Display(state [][]uint) string```
 Display allows us to see our puzzle in a nice grid. Simpily pass it the uint 2d slice, and it will return a nice string.
 See all other sections for display examples!

## PuzzleError
``` 
type PuzzleError struct {
	msg string
	Row uint
	Col uint
}
```
PuzzleError is a type that allows beter error reporting of puzzle specific errors. On a puzzle error, you can see the row/col that failed. In the below example, we have taken our hint puzzle and added a failure to the bottom left corner:
```
	puzzle := Puzzle{
		{0, 0, 0, 6, 0, 4, 7, 0, 0},
		{7, 0, 6, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 8, 0},
		{0, 7, 0, 0, 2, 0, 0, 9, 3},
		{8, 0, 0, 0, 0, 0, 0, 0, 5},
		{4, 3, 0, 0, 1, 0, 0, 7, 0},
		{0, 5, 0, 2, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 2, 0, 8},
		{3, 0, 2, 3, 0, 1, 0, 0, 0},
	}

	fmt.Printf("Pre-hint puzzle:\n%s\n", solve.Display(puzzle))
	hintedPuzzle, row, col, err := solve.Hint(puzzle)
	if pe, ok := err.(*solve.PuzzleError); ok {
		fmt.Printf("Puzzled! At (%d,%d): %s", pe.Row, pe.Col, pe)
	} else {
		fmt.Printf("Getting a hint, post hint puzzle(%d,%d):\n%s\n", row, col, solve.Display(hintedPuzzle))
	}

```
will return this:

```
Pre-hint puzzle:
  0  0  0 |  6  0  4 |  7  0  0
  7  0  6 |  0  0  0 |  0  0  9
  0  0  0 |  0  0  5 |  0  8  0
---------------------------------
  0  7  0 |  0  2  0 |  0  9  3
  8  0  0 |  0  0  0 |  0  0  5
  4  3  0 |  0  1  0 |  0  7  0
---------------------------------
  0  5  0 |  2  0  0 |  0  0  0
  3  0  0 |  0  0  0 |  2  0  8
  3  0  2 |  3  0  1 |  0  0  0

Puzzled! At (7,0): Contradiction at (7,0) (value 3). Unable to resolve.Getting a hint, post hint puzzle(0,0):

```
