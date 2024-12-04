package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// print inputs
	inputs := loadInputs("inputs.txt")

	// print first part solution
	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs [][]string) int {
	return countMatches(inputs, "XMAS")
}

func secondPart(inputs [][]string) int {
	return countXMASShapedMatches(inputs)
}

func loadInputs(filename string) (inputs [][]string) {
	// It should load the inputs from the given file.

	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	fileContent := string(data)

	// Insert each line as an array of strings
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimSpace(line)

		inputs = append(inputs, strings.Split(line, ""))
	}

	return inputs
}

// function that returns the number of found who matches a string value in a 2D array
func countMatches(inputs [][]string, value string) int {
	count := 0

	valueLength := len(value)
	height := len(inputs)
	width := len(inputs[0])

	// Look for value : horizontally, vertically, and diagonally
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			allPossibleValues := [][]string{
				getRightHorizontalValues(inputs, i, j, valueLength),
				getLeftHorizontalValues(inputs, i, j, valueLength),
				getUpperVerticalValues(inputs, i, j, valueLength),
				getLowerVerticalValues(inputs, i, j, valueLength),
				getUpperLeftDiagonalValues(inputs, i, j, valueLength),
				getUpperRightDiagonalValues(inputs, i, j, valueLength),
				getLowerLeftDiagonalValues(inputs, i, j, valueLength),
				getLowerRightDiagonalValues(inputs, i, j, valueLength),
			}

			for _, values := range allPossibleValues {
				if isArrayEqualToStringArray(values, value) {
					count++
				}
			}
		}
	}

	return count
}

func isArrayEqualToStringArray(array []string, expected string) bool {
	if len(array) != len(expected) {
		return false
	}

	// Combine the array into a single string
	combined := strings.Join(array, "")

	return strings.ToUpper(combined) == strings.ToUpper(expected)
}

func getRightHorizontalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values to the right of the current position
	for i := 0; i < length; i++ {
		if column+i >= len(inputs[row]) {
			break
		}
		values = append(values, inputs[row][column+i])
	}

	return values
}

func getLeftHorizontalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values to the left of the current position
	for i := 0; i < length; i++ {
		if column-i < 0 {
			break
		}
		values = append(values, inputs[row][column-i])
	}

	return values
}

func getUpperVerticalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values above the current position
	for i := 0; i < length; i++ {
		if row-i < 0 {
			break
		}
		values = append(values, inputs[row-i][column])
	}

	return values
}

func getLowerVerticalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values below the current position
	for i := 0; i < length; i++ {
		if row+i >= len(inputs) {
			break
		}
		values = append(values, inputs[row+i][column])
	}

	return values
}

func getUpperLeftDiagonalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values in the upper left diagonal
	for i := 0; i < length; i++ {
		if row-i < 0 || column-i < 0 {
			break
		}
		values = append(values, inputs[row-i][column-i])
	}

	return values
}

func getUpperRightDiagonalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values in the upper right diagonal
	for i := 0; i < length; i++ {
		if row-i < 0 || column+i >= len(inputs[row]) {
			break
		}
		values = append(values, inputs[row-i][column+i])
	}

	return values
}

func getLowerLeftDiagonalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	// Get the values in the lower left diagonal
	for i := 0; i < length; i++ {
		if row+i >= len(inputs) || column-i < 0 {
			break
		}
		values = append(values, inputs[row+i][column-i])
	}

	return values
}

func getLowerRightDiagonalValues(inputs [][]string, row int, column int, length int) []string {
	values := []string{}

	maxLength := len(inputs)
	maxWidth := len(inputs[row])

	// Get the values in the lower right diagonal
	for i := 0; i < length; i++ {
		//fmt.Println("row+i", row+i, "maxLength", maxLength, "column+i", column+i, "maxWidth", maxWidth)
		if row+i >= maxLength || column+i >= maxWidth {
			break
		}

		values = append(values, inputs[row+i][column+i])
	}

	return values
}

func countXMASShapedMatches(inputs [][]string) int {
	count := 0

	height := len(inputs)
	width := len(inputs[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if isXMASShaped(inputs, i, j) {
				count++
			}
		}
	}

	return count
}

func isXMASShaped(inputs [][]string, row int, column int) bool {
	// It should return true if the given position is the center of an XMAS shaped match
	// It should return false otherwise

	// valid XMAS shaped match, it can be also reversed
	/*
		M.S
		.A.
		M.S
	*/
	searchedValues := []string{"MAS", "SAM"}
	middleValue := "A"
	valueSize := len(searchedValues[0])
	maxHeight := len(inputs)
	maxWidth := len(inputs[row])

	if inputs[row][column] != middleValue {
		return false
	}

	topRow := row - 1
	bottomRow := row + 1
	leftColumn := column - 1
	rightColumn := column + 1

	// If row, column is at the edge, it can't be the center of an XMAS shaped match
	if topRow < 0 || leftColumn < 0 || rightColumn >= maxWidth || bottomRow >= maxHeight {
		return false
	}

	// Check if the current position is the center of an XMAS shaped match
	leftValues := getLowerLeftDiagonalValues(inputs, topRow, rightColumn, valueSize)
	rightValues := getLowerRightDiagonalValues(inputs, topRow, leftColumn, valueSize)

	// Check if the left and right values are equal to the searched values
	isLeftMatch := isArrayEqualToStringArray(leftValues, searchedValues[0]) || isArrayEqualToStringArray(leftValues, searchedValues[1])
	isRightMatch := isArrayEqualToStringArray(rightValues, searchedValues[0]) || isArrayEqualToStringArray(rightValues, searchedValues[1])

	if isLeftMatch && isRightMatch {
		return true
	}

	return false
}
