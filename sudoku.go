package main

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

var sudoku = Sudoku{
	{0, 0, 0, 6, 0, 0, 0, 5, 4},
	{0, 3, 0, 0, 0, 0, 0, 0, 0},
	{6, 0, 0, 1, 8, 0, 0, 3, 9},

	{3, 0, 0, 0, 0, 2, 0, 9, 0},
	{0, 9, 0, 0, 7, 0, 1, 0, 0},
	{2, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 4, 0, 0, 0, 0, 0, 0},
	{0, 0, 7, 0, 6, 0, 0, 0, 8},
	{8, 0, 0, 0, 0, 7, 2, 6, 0},
}

func main() {
	sudoku.print()

}

/*package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y float64
}

func Abs(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vector{3, 4}
	fmt.Println(Abs(v))
}
*/
