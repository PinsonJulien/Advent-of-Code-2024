package main

import (
	"reflect"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	area := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	topographicMap := TopographicMap{area: area}

	return ProblemInput{topographicMap}
}

func TestLoadInputs(t *testing.T) {
	inputs := loadInputs(DefaultTestInputFile)

	if !reflect.DeepEqual(inputs, getTestInputs()) {
		t.Errorf("Expected %v, got %v", getTestInputs(), inputs)
	}
}

func TestTopographicMapGetScore(t *testing.T) {
	inputs := getTestInputs()

	expected := 36
	result := inputs.topographicMap.getScore()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestTopographicMapGetTrails(t *testing.T) {
	inputs := getTestInputs()

	expected := 9
	result := len(inputs.topographicMap.getTrails())

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
