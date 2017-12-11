package solve

import "testing"

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
