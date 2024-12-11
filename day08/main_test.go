package main

import (
	"reflect"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	cityMap := [][]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "0", ".", ".", "."},
		{".", ".", ".", ".", ".", "0", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "0", ".", ".", ".", "."},
		{".", ".", ".", ".", "0", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "A", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "A", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "A", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	inputs := ProblemInput{cityMap: cityMap}

	return inputs
}

func TestLoadInputs(t *testing.T) {
	inputs := loadInputs(DefaultTestInputFile)

	if !reflect.DeepEqual(inputs, getTestInputs()) {
		t.Errorf("Expected %v, got %v", getTestInputs(), inputs)
	}
}
