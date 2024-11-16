package main

type Position struct {
	rowIndex int
	colIndex int
}

type Sudoku [9][9]int

func (s *Sudoku) print() {
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
func (s *Sudoku) isSafeToPlace(value int, position Position) bool {
	for i := 0; i < 9; i++ {
		if s[position.rowIndex][i] == value {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if s[i][position.colIndex] == value {
			return false
		}
	}

	blockStart := Position{rowIndex: position.rowIndex - position.rowIndex%3, colIndex: position.colIndex - position.colIndex%3}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i+blockStart.rowIndex][j+blockStart.colIndex] == value {
				return false
			}
		}

	}

	return true
}

// todo: aint gona work with a method
func (s *Sudoku) solve(rowIndex int, colIndex int) bool {
	// todo: why 8???, because it's 0 indexed
	// then why 9? already increased too much somewhere?
	// todo: check where is colIndex increased
	if rowIndex == 8 && colIndex == 9 {
		return true
	}

	// yeah, we let it increase too much
	if colIndex == 9 {
		rowIndex++
		colIndex = 0
	}

	// step to next if already has value
	if s[rowIndex][colIndex] != 0 {
		return s.solve(rowIndex, colIndex+1)
	}

	// try possible values
	for number := 0; number < 10; number++ {
		if s.isSafeToPlace(number, Position{rowIndex, colIndex}) {
			s[rowIndex][colIndex] = number

			//  incremented here as well
			if s.solve(rowIndex, colIndex+1) {
				return true
			}
		}

		// we set the value back to 0, prob this is why it can work with mutable array
		s[rowIndex][colIndex] = 0
	}

	return false
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

func main() {
	sudoku.print()

	sudoku.solve(0, 0)
	sudoku.print()
}
