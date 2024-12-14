package main

import (
	"fmt"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	// 2333133121414131402
	files := []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}

	return ProblemInput{disk: Disk{files: files}}
}

func TestLoadInputs(t *testing.T) {
	inputs := loadInputs(DefaultTestInputFile)

	expected := getTestInputs()

	if len(inputs.disk.files) != len(expected.disk.files) {
		t.Errorf("Expected %d files, got %d", len(expected.disk.files), len(inputs.disk.files))
	}

	for i, file := range inputs.disk.files {
		if file != expected.disk.files[i] {
			t.Errorf("Expected %d, got %d", expected.disk.files[i], file)
		}
	}
}

func TestGetChecksum(t *testing.T) {
	inputs := getTestInputs()

	expected := 1928
	result := inputs.disk.getChecksum()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCompact(t *testing.T) {
	inputs := getTestInputs()

	// 0099811188827773336446555566
	expected := []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6}

	result := inputs.disk.compact()

	if len(result) != len(expected) {
		t.Errorf("Expected %d files, got %d", len(expected), len(result))
	}

	for i, file := range result {
		if file != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], file)
		}
	}
}

func TestGetWholeFileChecksum(t *testing.T) {
	inputs := getTestInputs()

	expected := 2858
	result := inputs.disk.getWholeFileChecksum()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCompactWithWholeFileChecksum(t *testing.T) {
	inputs := getTestInputs()

	//00992111777.44.333....5555.6666.....8888..
	expected := []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, 0, 4, 4, 0, 3, 3, 3, 0, 0, 0, 0, 5, 5, 5, 5, 0, 6, 6, 6, 6, 0, 0, 0, 0, 0, 8, 8, 8, 8, 0, 0}

	result := inputs.disk.compactWithWholeFile()
	fmt.Println(result)
	if len(result) != len(expected) {
		t.Errorf("Expected %d files, got %d", len(expected), len(result))
	}

	for i, file := range result {
		if file != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], file)
		}
	}
}

/*
[{false 0 2}
{true 1 3}
{false 1 3}
{true 2 3}
{false 2 1}
{true 3 3}
{false 3 3}
{true 4 1}
{false 4 2}
{true 5 1}
{false 5 4}
{true 6 1}
{false 6 4}
{true 7 1}
{false 7 3}
{true 8 1}
{false 8 4}
{true 9 0}
{false 9 2}]
*/
