package main

import (
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func GetFirstColumn() []int {
	return []int{3, 4, 2, 1, 3, 3}
}

func GetSecondColumn() []int {
	return []int{4, 3, 5, 3, 9, 3}
}

func GetTestPairs() []Pair {
	return []Pair{
		{1, 3},
		{3, 2},
		{3, 3},
		{4, 3},
		{3, 5},
		{9, 4},
	}
}

func TestLoadInputs(t *testing.T) {

	// it should return 2 arrays of integers, which are :
	// 	[3, 4, 2, 1, 3, 3]
	//	[4, 3, 5, 3, 9, 3]
	// from the inputs.txt.example file.
	var ExpectedFirstColumn = GetFirstColumn()
	var ExpectedSecondColumn = GetSecondColumn()

	firstColumn, secondColumn := LoadInputs(DefaultTestInputFile)

	for i := 0; i < len(ExpectedFirstColumn); i++ {
		if firstColumn[i] != ExpectedFirstColumn[i] {
			t.Errorf("Expected firstColumn[%d] to be %d, but got %d", i, ExpectedFirstColumn[i], firstColumn[i])
		}
		if secondColumn[i] != ExpectedSecondColumn[i] {
			t.Errorf("Expected secondColumn[%d] to be %d, but got %d", i, ExpectedSecondColumn[i], secondColumn[i])
		}
	}
}

func TestGetPairs(t *testing.T) {
	/*
		it should return 6 pairs, which are :
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	*/
	// They have to be sorted, therefore the expected pairs are :
	// 1, 3
	// 2, 3
	// 3, 3
	// 3, 4
	// 3, 5
	// 4, 9

	var ExpectedPairs = GetTestPairs()

	firstColumn := GetFirstColumn()
	secondColumn := GetSecondColumn()
	pairs := GetPairs(firstColumn, secondColumn)

	for i := 0; i < len(ExpectedPairs); i++ {
		if pairs[i] != ExpectedPairs[i] {
			t.Errorf("Expected pairs[%d] to be %v, but got %v", i, ExpectedPairs[i], pairs[i])
		}
	}
}

func TestPair_Distance(t *testing.T) {
	// it should return 2, 1, 0, 1, 2, 5
	var ExpectedDistances = []int{2, 1, 0, 1, 2, 5}

	pairs := GetTestPairs()

	for i := 0; i < len(ExpectedDistances); i++ {
		if pairs[i].Distance() != ExpectedDistances[i] {
			t.Errorf("Expected pairs[%d].Distance() to be %d, but got %d", i, ExpectedDistances[i], pairs[i].Distance())
		}
	}
}

func TestCalculateSumOfDistances(t *testing.T) {
	// it should return 11
	var ExpectedSum = 11

	pairs := GetTestPairs()

	if CalculateSumOfDistances(pairs) != ExpectedSum {
		t.Errorf("Expected CalculateSumOfDistances to be %d, but got %d", ExpectedSum, CalculateSumOfDistances(pairs))
	}
}

func TestCountOccurrences(t *testing.T) {

	TestValues := func(value int, column []int, expected int) {
		if CountOccurrences(value, column) != expected {
			t.Errorf("Expected CountOccurrences to be %d, but got %d", expected, CountOccurrences(value, column))
		}
	}

	// it should return 3
	var ExpectedCount = 3

	column := []int{4, 3, 5, 3, 9, 3}
	value := 3

	TestValues(value, column, ExpectedCount)

	// it should return 1
	ExpectedCount = 1
	value = 4

	TestValues(value, column, ExpectedCount)

	// it should return 0
	ExpectedCount = 0
	value = 2

	TestValues(value, column, ExpectedCount)

	// it should return 0
	ExpectedCount = 0
	value = 1

	TestValues(value, column, ExpectedCount)
}

func TestCalculateSimilarityScore(t *testing.T) {
	// it should return 31
	var ExpectedScore = 31

	firstColumn := []int{3, 4, 2, 1, 3, 3}
	secondColumn := []int{4, 3, 5, 3, 9, 3}

	if CalculateSimilarityScore(firstColumn, secondColumn) != ExpectedScore {
		t.Errorf("Expected CalculateSimilarityScore to be %d, but got %d", ExpectedScore, CalculateSimilarityScore(firstColumn, secondColumn))
	}
}
