# sudoKu
A library to help solve sudoKu puzzles, written in Go.


## Getting Started
Simpily just import this module, then you are ready to start using methods!


## Exposed functions

### SolvePuzzle
 ```SolvePuzzle(puzzle [][]uint) (solvedPuzzle [][]uint,err error)```
SolvePuzzle is your main call into sudoKu! it takes a 2d slice of uints and returns a solved 2d slice of uints, or an error. 
<example>

### Hint
 ```Hint(puzzle [][]uint) (updatedPuzzle [][]uint, row uint, col uint, err error)```
Hint takes a puzzle, and returns and updated puzzle and the row/col values + an error. This error can be a PuzzleErorr!
<example>

### Display
 ```Display(state [][]uint) string```
 Display allows us to see our puzzle in a nice grid. Simpily pass it the uint 2d slice, and it will return a nice string.
 See all other sections for display examples!

## Exposed Types

### PuzzleError
``` 
type PuzzleError struct {
	msg string
	Row uint
	Col uint
}
```
PuzzleError is a type that allows beter error reporting of puzzle specific errors. On a puzzle error, you can see the row/col that failed.
<example>
