package main

import (
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

// It should return an array of array of strings, which are :
func getTestInputs() (inputs [][]string) {
	return [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}
}

func TestLoadInputs(t *testing.T) {

	// It should load the inputs from the inputs.txt.example file.
	/*
		MMMSXXMASM
		MSAMXMSMSA
		AMXSXMAAMM
		MSAMASMSMX
		XMASAMXAMM
		XXAMMXXAMA
		SMSMSASXSS
		SAXAMASAAA
		MAMMMXMMMM
		MXMXAXMASX
	*/
	// It should an array of array of characters, which are :
	// 	[MMMSXXMASM]
	//	[MSAMXMSMSA]
	//	[AMXSXMAAMM]
	//	[MSAMASMSMX]
	//	[XMASAMXAMM]
	//	[XXAMMXXAMA]
	//	[SMSMSASXSS]
	//	[SAXAMASAAA]
	//	[MAMMMXMMMM]
	//	[MXMXAXMASX]

	var expectedInputs = getTestInputs()

	inputs := loadInputs(DefaultTestInputFile)

	for i := 0; i < len(expectedInputs); i++ {
		for j := 0; j < len(expectedInputs[i]); j++ {
			if inputs[i][j] != expectedInputs[i][j] {
				t.Errorf("Expected inputs[%d][%d] to be %s, but got %s", i, j, expectedInputs[i][j], inputs[i][j])
			}
		}
	}
}

func TestCountMatches(t *testing.T) {
	// Search for "XMAS" should return 18.

	expectedResult := 18
	inputs := getTestInputs()

	result := countMatches(inputs, "XMAS")

	if result != expectedResult {
		t.Errorf("Expected result to be %d, but got %d", expectedResult, result)
	}
}

func TestGetRightHorizontalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getRightHorizontalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{"M", "M", "M", "S"}
	testValues(inputs, row, column, length, expectedValues)

	// for 3rd row
	row = 2
	column = 9
	expectedValues = []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 8
	expectedValues = []string{"M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 7
	expectedValues = []string{"A", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 6
	expectedValues = []string{"A", "A", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 5
	expectedValues = []string{"M", "A", "A", "M"}
}

func TestGetLeftHorizontalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getLeftHorizontalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 9
	expectedValues = []string{"M", "M", "A", "A"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 8
	expectedValues = []string{"M", "A", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 0
	expectedValues = []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 1
	expectedValues = []string{"M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 2
	expectedValues = []string{"M", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 3
	expectedValues = []string{"S", "M", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestGetUpperVerticalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getUpperVerticalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 1
	column = 0
	expectedValues = []string{"M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 0
	expectedValues = []string{"A", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 3
	column = 0
	expectedValues = []string{"M", "A", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 4
	column = 0
	expectedValues = []string{"X", "M", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 5
	column = 5
	expectedValues = []string{"X", "M", "S", "M"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestGetLowerVerticalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getLowerVerticalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{"M", "M", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 1
	column = 0
	expectedValues = []string{"M", "A", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 9
	column = 3
	expectedValues = []string{"X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 8
	column = 3
	expectedValues = []string{"M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 7
	column = 3
	expectedValues = []string{"A", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 6
	column = 3
	expectedValues = []string{"M", "A", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestGetUpperLeftDiagonalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getUpperLeftDiagonalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 1
	column = 1
	expectedValues = []string{"S", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 2
	expectedValues = []string{"X", "S", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 3
	column = 3
	expectedValues = []string{"M", "X", "S", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 4
	column = 4
	expectedValues = []string{"A", "M", "X", "S"}
	testValues(inputs, row, column, length, expectedValues)

	row = 9
	column = 2
	expectedValues = []string{"M", "A", "S"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestGetUpperRightDiagonalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getUpperRightDiagonalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 0
	column := 0
	expectedValues := []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 1
	column = 1
	expectedValues = []string{"S", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 2
	column = 2
	expectedValues = []string{"X", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 3
	column = 3
	expectedValues = []string{"M", "X", "M", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 4
	column = 4
	expectedValues = []string{"A", "S", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 9
	column = 7
	expectedValues = []string{"A", "M", "A"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestGetLowerLeftDiagonalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getLowerLeftDiagonalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 9
	column := 0
	expectedValues := []string{"M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 8
	column = 1
	expectedValues = []string{"A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 7
	column = 2
	expectedValues = []string{"X", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 6
	column = 3
	expectedValues = []string{"M", "X", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 5
	column = 4
	expectedValues = []string{"M", "M", "X", "A"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 2
	expectedValues = []string{"M", "S", "A"}
}

func TestGetLowerRightDiagonalValues(t *testing.T) {
	testValues := func(inputs [][]string, row, column, length int, expectedValues []string) {
		values := getLowerRightDiagonalValues(inputs, row, column, length)

		for i := 0; i < len(expectedValues); i++ {
			if values[i] != expectedValues[i] {
				t.Errorf("Expected values[%d] to be %s, but got %s. (row: %d, col: %d)", i, expectedValues[i], values[i], row, column)
			}
		}
	}

	searchedValue := "XMAS"
	length := len(searchedValue)
	inputs := getTestInputs()

	row := 9
	column := 9
	expectedValues := []string{"X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 8
	column = 8
	expectedValues = []string{"M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 7
	column = 7
	expectedValues = []string{"A", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 6
	column = 6
	expectedValues = []string{"S", "A", "M", "X"}
	testValues(inputs, row, column, length, expectedValues)

	row = 5
	column = 5
	expectedValues = []string{"X", "S", "A", "M"}
	testValues(inputs, row, column, length, expectedValues)

	row = 0
	column = 7
	expectedValues = []string{"A", "S", "M"}
	testValues(inputs, row, column, length, expectedValues)
}

func TestCountXMASShapedMatches(t *testing.T) {
	// it should return 9, since there are 9 XMAS shaped matches in the test inputs
	expectedResult := 9
	inputs := getTestInputs()

	result := countXMASShapedMatches(inputs)

	if result != expectedResult {
		t.Errorf("Expected result to be %d, but got %d", expectedResult, result)
	}
}

func TestIsXMASShaped(t *testing.T) {
	/*
		Check the following XMAS shaped match in the test inputs :

		.M.S......
		..A..MSMS.
		.M.S.MAA..
		..A.ASMSM.
		.M.S.M....
		..........
		S.S.S.S.S.
		.A.A.A.A..
		M.M.M.M.M.
		..........
	*/

	// valid XMAS shaped match, it can be also reversed
	/*
		M.S
		.A.
		M.S
	*/

	testValues := func(inputs [][]string, row, column int, expected bool) {
		result := isXMASShaped(inputs, row, column)

		if result != expected {
			t.Errorf("Expected result to be %t, but got %t. (row: %d, col: %d)", expected, result, row, column)
		}
	}

	inputs := getTestInputs()

	// Check borders
	row := 0
	column := 0
	testValues(inputs, row, column, false)

	row = 0
	column = 4
	testValues(inputs, row, column, false)

	row = 1
	column = 0
	testValues(inputs, row, column, false)

	row = 6
	column = 0
	testValues(inputs, row, column, false)

	// Check valid XMAS shaped match
	row = 1
	column = 2
	testValues(inputs, row, column, true)

	row = 2
	column = 6
	testValues(inputs, row, column, true)

	row = 2
	column = 7
	testValues(inputs, row, column, true)

	row = 3
	column = 2
	testValues(inputs, row, column, true)

	row = 3
	column = 4
	testValues(inputs, row, column, true)

	row = 7
	column = 1
	testValues(inputs, row, column, true)

	row = 7
	column = 3
	testValues(inputs, row, column, true)

	row = 7
	column = 5
	testValues(inputs, row, column, true)

	row = 7
	column = 7
	testValues(inputs, row, column, true)
	
}
