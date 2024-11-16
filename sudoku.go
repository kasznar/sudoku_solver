package main

import "fmt"

type Sudoku [][]int

func (s Sudoku) print() {
	for rangeIndex, row := range s {
		print("|")
		for numberIndex, number := range row {
			print(number)
			print(" ")
			if numberIndex == 2 || numberIndex == 5 {
				print("|")
			}
		}
		print("|")

		if rangeIndex == 2 || rangeIndex == 5 {
			println()
			for i := 0; i < 22; i++ {
				print("-")
			}
		}

		println()
	}
}

// todo: what if we already have value there?
func (s Sudoku) isSafeToPlace(rowIndex int, colIndex int, value int) bool {
	for i := 0; i < 9; i++ {
		if s[rowIndex][i] == value {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if s[i][colIndex] == value {
			return false
		}
	}

	blockStartRowIndex := rowIndex - rowIndex%3
	blockStartColIndex := colIndex - colIndex%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i+blockStartRowIndex][j+blockStartColIndex] == value {
				return false
			}
		}

	}

	return true
}

var sudoku = Sudoku{
	{3, 0, 6, 5, 0, 8, 4, 0, 0},
	{5, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 8, 7, 0, 0, 0, 0, 3, 1},
	{0, 0, 3, 0, 1, 0, 0, 8, 0},
	{9, 0, 0, 8, 6, 3, 0, 0, 5},
	{0, 5, 0, 0, 9, 0, 6, 0, 0},
	{1, 3, 0, 0, 0, 0, 2, 5, 0},
	{0, 0, 0, 0, 0, 0, 0, 7, 4},
	{0, 0, 5, 2, 0, 6, 3, 0, 0},
}

var solution = Sudoku{
	{3, 1, 6, 5, 7, 8, 4, 9, 2},
	{5, 2, 9, 1, 3, 4, 7, 6, 8},
	{4, 8, 7, 6, 2, 9, 5, 3, 1},
	{2, 6, 3, 4, 1, 5, 9, 8, 7},
	{9, 7, 4, 8, 6, 3, 1, 2, 5},
	{8, 5, 1, 7, 9, 2, 6, 4, 3},
	{1, 3, 8, 9, 4, 7, 2, 5, 6},
	{6, 9, 2, 3, 5, 1, 8, 7, 4},
	{7, 4, 5, 2, 8, 6, 3, 1, 9},
}

func testSafetyCheck(rowIndex int, colIndex int, value int) {
	print("It is ")

	if sudoku.isSafeToPlace(rowIndex, colIndex, value) {
		print("Safe")
	} else {
		print("Unsafe")
	}

	fmt.Printf(" to place %v at [%v][%v] \n", value, rowIndex, colIndex)
}

func main() {
	sudoku.print()
	testSafetyCheck(0, 1, 1)
	testSafetyCheck(0, 1, 3)
	testSafetyCheck(0, 1, 8)
}
