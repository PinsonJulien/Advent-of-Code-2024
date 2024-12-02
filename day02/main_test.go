package main

import (
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func GetTestReports() [][]int {
	return [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
}

func TestLoadInputs(t *testing.T) {
	// it should return all reports from the inputs.txt.example file.
	expectedReports := GetTestReports()

	reports := LoadInputs(DefaultTestInputFile)

	for i := 0; i < len(expectedReports); i++ {
		for j := 0; j < len(expectedReports[i]); j++ {
			if reports[i][j] != expectedReports[i][j] {
				t.Errorf("Expected reports[%d][%d] to be %d, but got %d", i, j, expectedReports[i][j], reports[i][j])
			}
		}
	}
}

func TestCountSafeReports(t *testing.T) {
	// it should return 3, since there are 3 safe reports in the test reports
	reports := GetTestReports()
	expectedCount := 2

	count := CountSafeReports(reports)

	if count != expectedCount {
		t.Errorf("Expected count to be %d, but got %d", expectedCount, count)
	}
}

func TestCountSafeReportsWithProblemDampener(t *testing.T) {
	// it should return 4, since there are 4 safe reports in the test reports
	reports := GetTestReports()
	expectedCount := 4

	count := CountSafeReportsWithDampener(reports)

	if count != expectedCount {
		t.Errorf("Expected count to be %d, but got %d", expectedCount, count)
	}
}

func TestIsReportSafe(t *testing.T) {
	testValues := func(report []int, expectedResult bool) {
		result := IsReportSafe(report)
		if result != expectedResult {
			t.Errorf("Expected %v, but got %v. Report: %v", expectedResult, result, report)
		}
	}

	// Safe because all decreasing by 1 or 2
	report := []int{7, 6, 4, 2, 1}
	expectedResult := true
	testValues(report, expectedResult)

	// unsafe because 2 and 7 are not within 1-3 range
	report = []int{1, 2, 7, 8, 9}
	expectedResult = false
	testValues(report, expectedResult)

	// unsafe because 6 and 2 are not within 1-3 range
	report = []int{9, 7, 6, 2, 1}
	expectedResult = false
	testValues(report, expectedResult)

	// unsafe because 1 and 3 are increasing but 3 and 2 are decreasing
	report = []int{1, 3, 2, 4, 5}
	expectedResult = false
	testValues(report, expectedResult)

	// unsafe because 4 and 4 are not an increase or decrease
	report = []int{8, 6, 4, 4, 1}
	expectedResult = false
	testValues(report, expectedResult)

	// safe because all increasing by 1, 2 or 3
	report = []int{1, 3, 6, 7, 9}
	expectedResult = true
	testValues(report, expectedResult)

	// edge case: 34 and 35 are increasing by 1, but 35 and 34 are decreasing by 1
	// therefore, it should return false
	report = []int{34, 35, 34, 31, 28}
	expectedResult = false
	testValues(report, expectedResult)
}

func TestRemoveIndexFromArray(t *testing.T) {

	testValues := func(array []int, index int, expectedArray []int) {
		result := RemoveIndexFromArray(array, index)
		for i := 0; i < len(expectedArray); i++ {
			if result[i] != expectedArray[i] {
				t.Errorf("Expected %v, but got %v", expectedArray, result)
			}
		}
	}

	array := []int{1, 2, 3, 4, 5}
	index := 2
	expectedArray := []int{1, 2, 4, 5}
	testValues(array, index, expectedArray)

	array = []int{1, 2, 3, 4, 5}
	index = 0
	expectedArray = []int{2, 3, 4, 5}
	testValues(array, index, expectedArray)

	array = []int{1, 2, 3, 4, 5}
	index = 4
	expectedArray = []int{1, 2, 3, 4}
	testValues(array, index, expectedArray)
}

func TestRemoveIndexFromArray_ShouldNotModifyOriginalArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	index := 2
	expectedArray := []int{1, 2, 3, 4, 5}
	RemoveIndexFromArray(array, index)
	for i := 0; i < len(expectedArray); i++ {
		if array[i] != expectedArray[i] {
			t.Errorf("Expected %v, but got %v", expectedArray, array)
		}
	}
}
