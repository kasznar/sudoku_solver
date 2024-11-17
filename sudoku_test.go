package main

import (
	"reflect"
	"testing"
)

func TestInsertSafetyCheck(t *testing.T) {
	input := []struct {
		rowI     int
		colI     int
		value    int
		expected bool
	}{
		{rowI: 0, colI: 1, value: 1, expected: true},
		{rowI: 0, colI: 1, value: 3, expected: false},
		{rowI: 0, colI: 1, value: 8, expected: false},
		{rowI: 0, colI: 2, value: 4, expected: false},
	}

	for _, data := range input {
		result := sudoku.isSafeToPlace(data.value, Position{colIndex: data.colI, rowIndex: data.rowI})
		if result != data.expected {
			t.Fatal("Check Not right ")
		}
	}
}

func TestSolution(t *testing.T) {
	testValue := sudoku.copy()

	testValue.solve(Position{0, 0})

	println("Result")
	testValue.print()

	println("Solution")
	solution.print()

	if !reflect.DeepEqual(solution, testValue) {
		t.Fatal("Wrong solution")
	}
}

func TestCopy(t *testing.T) {
	copied := solution.copy()

	if !reflect.DeepEqual(copied, solution) {
		t.Fatal("Copy doesn't match")
	}
}

func TestStepThrough(t *testing.T) {
	testValue := solution.copy()

	testValue.solve(Position{0, 0})

	if !reflect.DeepEqual(solution, testValue) {
		t.Fatal("Wrong solution")
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
/*func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}*/

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
/*func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}*/
