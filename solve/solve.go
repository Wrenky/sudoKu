/*

Follows the method from Peter Norvig's essay: http://norvig.com/sudoku.html
Uses constraint propagation and depth first search.
*/

package solve

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//The name of a square, which is the discrete entity that needs to be filled
type Square string

//The possible values that remain for each Square
type SquareOptions map[Square]string

// These are constants set in init, but needed as defaults in every puzzle
var digits, rows, cols = "123456789", "ABCDEFGHI", "123456789"

//All squares (A1, B1, C1... A2, B2, C2)
var squares []Square

// All regions- Rows, Columns and regional squares
var regions [][]Square

// Given a square, what are the squares we care about?
var linkedSquares map[Square][]Square

// Given a square, what are the regions that need to be solved?
var linkedRegions map[Square][][]Square

// Generate all of our constants on package initialiation!
func init() {

	// Helper function to combine strings together!
	cross := func(I string, J string) []Square {
		var out []Square
		for _, i := range I {
			for _, j := range J {
				out = append(out, Square(i)+Square(j))
			}
		}

		return out
	}

	// Use cross to create all of our squares
	squares = cross(rows, cols)

	// Now generate our regions
	regions := [][]Square{}
	rowBlocks := []string{"ABC", "DEF", "GHI"}
	rowCols := []string{"123", "456", "789"}
	for _, c := range cols {
		regions = append(regions, cross(rows, string(c)))
	}

	for _, r := range rows {
		regions = append(regions, cross(string(r), cols))
	}

	for _, rs := range rowBlocks {
		for _, cs := range rowCols {
			regions = append(regions, cross(rs, cs))
		}
	}
	linkedRegions = map[Square][][]Square{}
	for _, square := range squares {
	Square:
		for _, v := range regions {
			for _, linkedRegionsquare := range v {
				//If the square is contained within the regions, then add this regions belongs to this square's unit map
				if square == linkedRegionsquare {
					linkedRegions[square] = append(linkedRegions[square], v)
					continue Square
				}
			}
		}
	}

	linkedSquares = map[Square][]Square{}
	for square, regions := range linkedRegions {
		for _, unit := range regions {
		NextlinkedRegionsquare:
			for _, linkedRegionsquare := range unit {
				if linkedRegionsquare != square {
					for _, ps := range linkedSquares[square] {
						if ps == linkedRegionsquare {
							continue NextlinkedRegionsquare
						}
					}

					linkedSquares[square] = append(linkedSquares[square], linkedRegionsquare)
				}
			}
		}
	}
}

func parseToPuzzle(puzzle [][]uint) (SquareOptions, error) {

	var err error
	values := SquareOptions{}
	grid := SquareOptions{}
	// Populate grid with the values from puzzle!
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			value := strconv.FormatUint(uint64(puzzle[row][col]), 10)
			squareDex := ((row) * len(puzzle)) + (col)
			if value == "0" {
				grid[squares[squareDex]] = "."
			} else {
				grid[squares[squareDex]] = value
			}
		}
	}

	// Popuate values.
	for _, square := range squares {
		values[square] = digits
	}

	// Get grid filled with values
	for s, d := range grid {
		//Test each allowed digit
		for _, xd := range digits {
			//If the GridMap square's value is an allowed digit (not just a blank placeholder)
			if d == string(xd) {
				//Try to assign that value
				values, err = assign(values, s, d)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return values, nil
}

//Assign square s to digit d, and propagate.
func assign(state SquareOptions, s Square, d string) (SquareOptions, error) {

	//If this square has no possible state, then we are done.
	if len(state[s]) < 1 {
		return nil, errors.New(fmt.Sprintf("state[%s] has no possible digits left.", s))
	}

	otherSquareOptions := ``
	for _, v := range state[s] {
		//This value, d, is by definition not one of the "other" state
		if string(v) == d {
			continue
		}

		otherSquareOptions = otherSquareOptions + string(v)
	}

	// Eliminate values
	if len(otherSquareOptions) > 0 {
		for _, d2 := range otherSquareOptions {
			if _, err := eliminate(state, s, string(d2)); err != nil {
				return nil, err
			}
		}
	}
	return state, nil
}

//Eliminate digit d from the list of possible values at square s (values[s]) and propagate.
func eliminate(values SquareOptions, s Square, d string) (SquareOptions, error) {
	err := error(nil)

	dInSquareOptionsS := false
	for _, val := range values[s] {
		if string(val) == d {
			dInSquareOptionsS = true
			break
		}
	}

	// Then its already been eliminated!
	if !dInSquareOptionsS {
		return values, nil
	}

	values[s] = strings.Replace(values[s], d, "", -1)

	//If we have no more options, then we picked wrong. Fail.
	if len(values[s]) == 0 {
		return nil, errors.New(fmt.Sprintf("Cannot eliminate %s from values[%s] because now values[%s] has no valid potential digits (valid: %s).", d, s, s, values[s]))
	} else if len(values[s]) == 1 {
		// if only one thing is left, then this is the solution! (For this branch)
		d2 := values[s]
		for _, s2 := range linkedSquares[s] {
			// Propagate this constraint
			if _, err = eliminate(values, s2, d2); err != nil {
				return nil, err
			}
		}
	}

	for _, u := range linkedRegions[s] {
		//Iterate over all of the other squares in this square's linkedRegions.
		// For this unit, grab all of the squares that can accept `d` (enumerate them in dPlaces)
		dPlaces := []Square{}
		//For every square in the unit
		for _, s2 := range u {
			//For every digit that this square can accept
			for _, d2 := range values[s2] {
				//If d is in that square's digit list
				if d == string(d2) {
					//Then dPlaces includes this square
					dPlaces = append(dPlaces, s2)
					break
				}
			}
		}

		if len(dPlaces) == 0 {
			//No place to put d; contradiction
			return values, errors.New(fmt.Sprintf("There is no place in unit %s to put %s.", u, d))
		} else if len(dPlaces) == 1 {
			//D must go into dPlaces[0]
			_, err = assign(values, dPlaces[0], d)
			if err != nil {
				//If it can't, you'll have to backtrack (or you have an insoluble puzzle, depending on nesting depth)
				return nil, err
			}
		}
	}

	return values, nil
}
func SolvePuzzle(puzzle [][]uint) ([][]uint, error) {
	solved := [][]uint{}
	values, err := parseToPuzzle(puzzle)
	if err != nil {
		return solved, err
	}
	var results SquareOptions
	results, err = search(values, nil)
	if err != nil {
		return solved, err
	}

	solved = convertMapToSlice(results)

	return solved, nil
}

// Convert the solved map into slice form
func convertMapToSlice(puzzleMap SquareOptions) [][]uint {
	translation := make([][]uint, 9)
	for i := range translation {
		translation[i] = make([]uint, 9)
	}
	var combined string
	//This will loop A1, A2, A3... B.. C..
	//row+col should be A1, B2 while i,col is index
	for i, row := range rows {
		for col := range cols {
			var x, y = uint(i), uint(col)
			col = col + 1
			combined = string(row) + strconv.Itoa(int(col))
			value, _ := strconv.Atoi(puzzleMap[Square(combined)])
			translation[x][y] = uint(value)
		}
	}
	return translation
}

// Perform Depth first search on the remaining options in the map
func search(state SquareOptions, err error) (SquareOptions, error) {

	if err != nil {
		return state, err
	}

	// Check if we are finished
	solved := true
	for _, s := range squares {
		if len(state[s]) != 1 {
			solved = false
		}
	}
	if solved {
		return state, nil
	}

	// Here, we need to solve more. Take each square with more than one option, then
	// Assign a value and propagate until we fail or the puzzle is solved. Pick a square with
	// the least amount of possiblities to fail/solve quicker

	min, sq := len(digits)+1, Square("")
	for _, s := range squares {
		if len(state[s]) < min && len(state[s]) > 1 {
			min = len(state[s])
			sq = s

			if min == 2 {
				break
			}
		}
	}

	// Assign the value, and continue to search!
	for _, d := range state[sq] {
		vCloned, err := assign(cloneSquareOptions(state), sq, string(d))
		if err != nil {
			continue
		}
		vCloned, err = search(vCloned, err)
		if err == nil {
			return vCloned, nil
		}
	}

	return nil, errors.New("Your depth-first search failed on this branch.")
}

//Clone the state
func cloneSquareOptions(state SquareOptions) SquareOptions {
	cpySquareOptions := make(SquareOptions, len(state))
	for k, v := range state {
		cpySquareOptions[k] = v
	}

	return cpySquareOptions
}

// Displace the slice as a puzzle in ticktaktoe like form
func Display(state [][]uint) string {
	var out string
	for i := 0; i < len(state); i++ {
		if i == 3 || i == 6 {
			out = out + " -----------------------\n"
		}
		for j := 0; j < len(state[0]); j++ {
			if j == 3 || j == 6 {
				out = out + " | "
			}
			out = out + " " + strconv.Itoa(int(state[uint(i)][uint(j)]))
		}
		out = out + "\n"
	}
	return out
}