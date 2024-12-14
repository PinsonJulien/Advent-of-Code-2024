package main

import (
	"image"
	"reflect"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	inputs := NewProblemInput()
	size := 12
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			inputs.addPoint(image.Pt(x, y))
		}
	}

	// Add frequencies
	inputs.addFrequency('0', image.Pt(8, 1))
	inputs.addFrequency('0', image.Pt(5, 2))
	inputs.addFrequency('0', image.Pt(7, 3))
	inputs.addFrequency('0', image.Pt(4, 4))
	inputs.addFrequency('A', image.Pt(6, 5))
	inputs.addFrequency('A', image.Pt(8, 8))
	inputs.addFrequency('A', image.Pt(9, 9))

	return inputs
}

func TestLoadInputs(t *testing.T) {
	inputs := loadInputs(DefaultTestInputFile)

	if !reflect.DeepEqual(inputs, getTestInputs()) {
		t.Errorf("Expected %v, got %v", getTestInputs(), inputs)
	}
}

func TestCountAllAntiNodes(t *testing.T) {
	inputs := getTestInputs()

	expected := 14
	result := countAllAntiNodes(inputs)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetAllAntiNodes(t *testing.T) {
	inputs := getTestInputs()

	/*
		......#....#
		...#....0...
		....#0....#.
		..#....0....
		....0....#..
		.#....A.....
		...#........
		#......#....
		........A...
		.........A..
		..........#.
		..........#.
	*/

	expected := []image.Point{
		image.Pt(0, 7),
		image.Pt(1, 5),
		image.Pt(2, 3),
		image.Pt(3, 1),
		image.Pt(3, 6),
		image.Pt(4, 2),
		image.Pt(6, 0),
		image.Pt(6, 5),
		image.Pt(7, 7),
		image.Pt(9, 4),
		image.Pt(10, 2),
		image.Pt(10, 10),
		image.Pt(10, 11),
		image.Pt(11, 0),
	}

	result := getAllAntiNodes(inputs)

	// check length
	if len(result) != len(expected) {
		t.Errorf("Expected %d points, got %d", len(expected), len(result))
	}

	// Check if the result contains all the expected points, regardless of the order
	expectedMap := map[image.Point]struct{}{}
	for _, point := range expected {
		expectedMap[point] = struct{}{}
	}

	for _, point := range result {
		if _, ok := expectedMap[point]; !ok {
			t.Errorf("Unexpected point %v", point)
		}
	}
}

func TestCountAllAntiNodesWithResonantHarmonics(t *testing.T) {
	inputs := getTestInputs()

	expected := 34
	result := countAllAntiNodesWithResonantHarmonics(inputs)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetAllAntiNodesWithResonantHarmonics(t *testing.T) {
	inputs := getTestInputs()

	/*
		##....#....#
		.#.#....0...
		..#.#0....#.
		..##...0....
		....0....#..
		.#...#A....#
		...#..#.....
		#....#.#....
		..#.....A...
		....#....A..
		.#........#.
		...#......##
	*/
	
	expected := []image.Point{
		image.Pt(0, 0),
		image.Pt(0, 7),
		image.Pt(1, 0),
		image.Pt(1, 1),
		image.Pt(1, 5),
		image.Pt(1, 10),
		image.Pt(2, 2),
		image.Pt(2, 3),
		image.Pt(2, 8),
		image.Pt(3, 1),
		image.Pt(3, 3),
		image.Pt(3, 6),
		image.Pt(3, 11),
		image.Pt(4, 2),
		image.Pt(4, 4),
		image.Pt(4, 9),
		image.Pt(5, 2),
		image.Pt(5, 5),
		image.Pt(5, 7),
		image.Pt(6, 0),
		image.Pt(6, 5),
		image.Pt(6, 6),
		image.Pt(7, 3),
		image.Pt(7, 7),
		image.Pt(8, 1),
		image.Pt(8, 8),
		image.Pt(9, 4),
		image.Pt(9, 9),
		image.Pt(10, 2),
		image.Pt(10, 10),
		image.Pt(10, 11),
		image.Pt(11, 0),
		image.Pt(11, 5),
		image.Pt(11, 11),
	}

	result := getAllAntiNodesWithResonantHarmonics(inputs)

	// check length
	if len(result) != len(expected) {
		t.Errorf("Expected %d points, got %d", len(expected), len(result))
	}

	// Check if the result contains all the expected points, regardless of the order
	expectedMap := map[image.Point]struct{}{}
	for _, point := range expected {
		expectedMap[point] = struct{}{}
	}

	for _, point := range result {
		if _, ok := expectedMap[point]; !ok {
			t.Errorf("Unexpected point %v", point)
		}
	}

}
