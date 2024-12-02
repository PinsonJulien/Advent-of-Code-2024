package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := LoadInputs("inputs.txt")

	// print first part solution
	fmt.Println("First part solution: ", FirstPart(reports))
	fmt.Println("Second part solution: ", SecondPart(reports))
}

func FirstPart(reports [][]int) int {
	return CountSafeReports(reports)
}

func SecondPart(reports [][]int) int {
	return CountSafeReportsWithDampener(reports)
}

func LoadInputs(filename string) (reports [][]int) {
	// reading the inputs file as a matrix of integers
	// each line is a row, and each value is a column, separated by a space
	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		var report []int
		for _, column := range columns {
			value, err := strconv.Atoi(column)
			if err != nil {
				os.Exit(1)
			}
			report = append(report, value)
		}
		reports = append(reports, report)
	}

	return reports
}

func CountSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if IsReportSafe(report) {
			count++
		}
	}
	return count
}

func CountSafeReportsWithDampener(reports [][]int) int {
	count := 0
	for _, report := range reports {
		// If the report is already safe, we don't need to remove any level
		if IsReportSafe(report) {
			count++
			continue
		}

		// If the report is not safe, we try to remove each level and check if the report is safe
		for i := 0; i < len(report); i++ {
			reportWithoutUnsafeLevel := RemoveIndexFromArray(report, i)
			if IsReportSafe(reportWithoutUnsafeLevel) {
				count++
				break
			}
		}
	}
	return count
}

func IsReportSafe(report []int) bool {
	// Returns the index of the first level that is not safe

	// A report is safe if :
	// The levels are either all increasing or all decreasing
	// Two adjacent levels differ by at least 1 and at most 3 ; it cannot be 0.

	isWithinAllowedRange := func(value int) bool {
		const MinDifference = 1
		const MaxDifference = 3

		// convert value to positive
		if value < 0 {
			value = -value
		}

		return value >= MinDifference && value <= MaxDifference
	}

	length := len(report)

	// First we check the trend.
	firstDistance := CalculateDistance(report[0], report[1])
	if !isWithinAllowedRange(firstDistance) {
		return false
	}

	isIncreasing := firstDistance > 0

	// Start from the second element
	for i := 1; i < length-1; i++ {
		distance := CalculateDistance(report[i], report[i+1])
		isCurrentIncreasing := distance > 0
		if !isWithinAllowedRange(distance) {
			return false
		}

		if isIncreasing != isCurrentIncreasing {
			return false
		}
	}

	return true
}

func CalculateDistance(a int, b int) int {
	return a - b
}

func RemoveIndexFromArray(array []int, index int) []int {
	newSlice := append([]int{}, array[:index]...)
	return append(newSlice, array[index+1:]...)
}
