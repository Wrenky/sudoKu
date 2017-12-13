package solve

import "testing"
import "fmt"

type Puzzle [][]uint

func TestSingleIterationPuzzle(t *testing.T) {
	t.Log("Testing single iteration puzzle")
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
	res, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Error("Failed during solving")
		t.Error(err)
	} else {
		t.Log("Solved:")
		output := Display(res)
		t.Log("\n" + output)
	}
}

func TestHardPuzzle(t *testing.T) {
	t.Log("Testing Hard puzzle")
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
	res, err := SolvePuzzle(hardPuzzle)
	if err != nil {
		t.Error("Failed during solving")
		t.Error(err)
	} else {
		t.Log("Solved:")
		output := Display(res)
		t.Log("\n" + output)
	}
}

func TestZeroPuzzle(t *testing.T) {
	t.Log("Testing all Zeros- Should find a random solution")
	hardPuzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	res, err := SolvePuzzle(hardPuzzle)
	if err != nil {
		t.Error("Failed during solving")
		t.Error(err)
	} else {
		t.Log("Solved:")
		output := Display(res)
		t.Log("\n" + output)
	}
}

func TestInvalidAmountOfRows(t *testing.T) {
	t.Log("Testing invalid puzzle")
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	_, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Correctly failed:", err)
	} else {
		t.Error("Didnt fail with an invalid row amount")
	}
}

func TestInvalidAmountOfColumns(t *testing.T) {
	t.Log("Testing invalid column amount")
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	_, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Correctly failed:", err)
	} else {
		t.Error("Didnt fail with an invalid row amount")
	}

	puzzle = Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	_, err = SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Correctly failed on the last column:", err)
	} else {
		t.Error("Didnt fail with an invalid row amount")
	}
}

func TestInvalidSquare(t *testing.T) {
	t.Log("Testing invalid square")
	puzzle := Puzzle{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 15, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	_, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Correctly failed:", err)
	} else {
		t.Error("Didnt fail with an invalid square :(")
	}
}

func TestReallyHard(t *testing.T) {
	t.Log("Testing invalid square")
	puzzle := Puzzle{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},

		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 3, 0},

		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	_, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Error("Failed to find a solution:", err)
	} else {
		t.Log("Found a solution!")
	}
}

func TestInvalid(t *testing.T) {
	t.Log("Testing a puzzle with no solution")
	puzzle := Puzzle{
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 3},

		{0, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},

		{0, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
	solution, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Correctly failed with the following error:", err)
	} else {
		out := Display(solution)
		t.Error("Somehow found a solution?\n", out)
	}
}

func TestAlreadySolved(t *testing.T) {
	t.Log("Testing a puzzle that has already been solved")
	puzzle := Puzzle{
		{2, 8, 7, 6, 1, 3, 9, 4, 5},
		{9, 4, 1, 7, 5, 2, 3, 6, 8},
		{6, 5, 3, 8, 4, 9, 2, 1, 7},

		{7, 6, 5, 9, 3, 8, 4, 2, 1},
		{1, 3, 2, 4, 7, 6, 8, 5, 9},
		{8, 9, 4, 1, 2, 5, 6, 7, 3},

		{4, 1, 9, 2, 8, 7, 5, 3, 6},
		{5, 7, 6, 3, 9, 4, 1, 8, 2},
		{3, 2, 8, 5, 6, 1, 7, 9, 4},
	}
	solution, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Error("Failed to solve a correct puzzle!", err)
	} else {
		out := Display(solution)
		t.Log("Correctly returned the source back to us:\n", out)
	}
}

func TestCompletedBadPuzzle(t *testing.T) {
	t.Log("Testing an invalid puzzle")
	puzzle := Puzzle{
		{2, 8, 7, 6, 1, 3, 9, 4, 5},
		{9, 4, 1, 7, 5, 2, 3, 6, 8},
		{6, 5, 3, 8, 4, 9, 2, 1, 7},

		{7, 6, 5, 9, 3, 8, 4, 2, 1},
		{1, 3, 2, 4, 7, 6, 8, 5, 9},
		{8, 9, 4, 1, 2, 5, 6, 7, 3},

		{4, 1, 9, 2, 8, 7, 5, 3, 6},
		{5, 7, 6, 3, 9, 4, 1, 8, 2},
		{3, 2, 8, 5, 6, 1, 9, 9, 4},
	}
	solution, err := SolvePuzzle(puzzle)
	if err != nil {
		t.Log("Failed (correctly) with the following error:", err)
	} else {
		out := Display(solution)
		t.Error("Passed a bad solution as correct:", out)
	}
}

func TestHint(t *testing.T) {
	t.Log("Testing a normal hint")
	puzzle := Puzzle{
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
	out := Display(puzzle)
	t.Log(fmt.Sprintf("Starting puzzle\n%s\n", out))
	updatedPuzzle, row, col, err := Hint(puzzle)
	if err != nil {
		t.Error("Unable to find a hint:", err)
	} else {
		out = Display(updatedPuzzle)
		t.Log(fmt.Sprintf("Hint recieved: %d @ (%d,%d) :\n%s\n", updatedPuzzle[row][col], row, col, out))
	}
}

func TestHintOnBadPuzzle(t *testing.T) {
	t.Log("Testing a normal hint")
	puzzle := Puzzle{
		{0, 0, 0, 6, 0, 4, 7, 0, 0},
		{9, 0, 6, 0, 0, 2, 0, 0, 9},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},

		{0, 7, 0, 0, 2, 0, 0, 9, 3},
		{8, 0, 0, 0, 0, 0, 0, 0, 5},
		{4, 3, 0, 0, 1, 0, 0, 7, 0},

		{0, 5, 0, 2, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 2, 0, 8},
		{0, 0, 2, 3, 0, 1, 0, 0, 0},
	}

	puz := Display(puzzle)
	t.Log(fmt.Sprintf("Pre hint puzzle: \n%s\n", puz))

	updatedPuzzle, row, col, err := Hint(puzzle)
	if err != nil {
		t.Log("Got an error getting hint (correctly, contradiction at row 1):", err)
	} else {
		out := Display(updatedPuzzle)
		t.Error(fmt.Sprintf("BAD Updated hint, at %d,%d:\n%s\n", row, col, out))
	}
}
