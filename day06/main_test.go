package main

import (
	"reflect"
	"slices"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	labMap := LabMap{
		area: [][]string{
			//0    1    2    3    4    5    6    7    8    9
			{".", ".", ".", ".", "#", ".", ".", ".", ".", "."}, // 0
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"}, // 1
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."}, // 2
			{".", ".", "#", ".", ".", ".", ".", ".", ".", "."}, // 3
			{".", ".", ".", ".", ".", ".", ".", "#", ".", "."}, // 4
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."}, // 5
			{".", "#", ".", ".", "^", ".", ".", ".", ".", "."}, // 6
			{".", ".", ".", ".", ".", ".", ".", ".", "#", "."}, // 7
			{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."}, // 8
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."}, // 9
		},
	}

	return ProblemInput{labMap}
}

func TestLoadInputs(t *testing.T) {
	expected := getTestInputs()

	inputs := loadInputs(DefaultTestInputFile)

	if !reflect.DeepEqual(inputs, expected) {
		t.Errorf("Expected %v but got %v", expected, inputs)
	}
}

func TestNewGuardian(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)

	expected := Position{4, 6}
	actual := guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	expectedDirection := North
	actualDirection := guardian.direction

	if actualDirection != expectedDirection {
		t.Errorf("Expected %v but got %v", expectedDirection, actualDirection)
	}

	expectedLeftTheLab := false
	actualLeftTheLab := guardian.leftTheLab

	if actualLeftTheLab != expectedLeftTheLab {
		t.Errorf("Expected %v but got %v", expectedLeftTheLab, actualLeftTheLab)
	}

	expectedVisitedPositions := []Position{
		{4, 6},
	}
	actualVisitedPositions := guardian.visitedPositions

	if !reflect.DeepEqual(actualVisitedPositions, expectedVisitedPositions) {
		t.Errorf("Expected %v but got %v", expectedVisitedPositions, actualVisitedPositions)
	}

	expectedLabMap := inputs.labMap
	actualLabMap := guardian.LabMap

	if !reflect.DeepEqual(actualLabMap, expectedLabMap) {
		t.Errorf("Expected %v but got %v", expectedLabMap, actualLabMap)
	}
}

func TestGuardianVisitTheLab(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)
	guardian.visitTheLab()

	expected := 41
	actual := guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestGuardianMoveForward(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)
	// Move up
	guardian.moveForward()

	expected := Position{4, 5}
	actual := guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	// Move up 5 times
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()

	// It should turn right
	expected = Position{6, 1}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move right 4 times
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()

	// next should go down
	guardian.moveForward()

	expected = Position{8, 4}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move down 5 times
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()
	guardian.moveForward()

	// next should go left
	guardian.moveForward()

	expected = Position{4, 6}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Simulate the guardian leaving the lab
	guardian.currentPosition = Position{
		x: 0,
		y: 0,
	}

	guardian.direction = North

	guardian.moveForward()

	if !guardian.leftTheLab {
		t.Errorf("Expected leftTheLab to be true")
	}
}

func TestGuardianGetNextPosition(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)

	// Move up
	guardian.moveForward()

	expected := Position{4, 5}
	actual := guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move right
	guardian.moveForward()

	expected = Position{4, 4}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move down
	guardian.moveForward()

	expected = Position{4, 3}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move left
	guardian.moveForward()

	expected = Position{4, 2}
	actual = guardian.currentPosition

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestGuardianTurn(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)

	guardian.turn()

	expected := East
	actual := guardian.direction

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	guardian.turn()

	expected = South
	actual = guardian.direction

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	guardian.turn()

	expected = West
	actual = guardian.direction

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	guardian.turn()

	expected = North
	actual = guardian.direction

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestGuardianAddPositionToSet(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)

	// Move up
	guardian.moveForward()

	expected := 2
	actual := guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move right
	guardian.moveForward()

	expected = 3
	actual = guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move down
	guardian.moveForward()

	expected = 4
	actual = guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Move left
	guardian.moveForward()

	expected = 5
	actual = guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestGuardianGetUniquePositionsVisited(t *testing.T) {
	inputs := getTestInputs()

	guardian := NewGuardian(inputs.labMap)

	// Move up
	guardian.moveForward()

	// Move right
	guardian.moveForward()

	// Move down
	guardian.moveForward()

	// Move left
	guardian.moveForward()

	expected := 5
	actual := guardian.getUniquePositionsVisited()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestLabMapGetPositionsWhichWouldCauseALoop(t *testing.T) {

	inputs := getTestInputs()

	LabMap := inputs.labMap

	expected := []Position{
		// left to guard starting position
		{3, 6},
		// bottom right quadrant
		{6, 7},
		// next to bottom right quadrant
		{7, 7},
		// bottom left corner
		{1, 8},
		// right to bottom left corner
		{3, 8},
		// bottom right corner
		{7, 9},
	}

	// should have 6 positions
	actual := LabMap.getPositionsWhichWouldCauseALoop()

	if len(actual) != len(expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// check regardless of order
	for _, position := range expected {
		if !slices.Contains(actual, position) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}

}

func TestLabMapIsPositionOutOfBounds(t *testing.T) {
	inputs := getTestInputs()

	labMap := inputs.labMap

	// Test out of bounds
	position := Position{10, 10}

	expected := true
	actual := labMap.isPositionOutOfBounds(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test in bounds
	position = Position{5, 5}

	expected = false
	actual = labMap.isPositionOutOfBounds(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestLabMapIsPositionWall(t *testing.T) {
	inputs := getTestInputs()

	labMap := inputs.labMap

	// Test wall
	position := Position{2, 3}

	expected := true
	actual := labMap.isPositionWall(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test empty
	position = Position{2, 4}

	expected = false
	actual = labMap.isPositionWall(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestLabMapGetCellAtPosition(t *testing.T) {
	inputs := getTestInputs()

	labMap := inputs.labMap

	// Test wall
	position := Position{2, 3}

	expected := "#"
	actual := labMap.getCellAtPosition(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test empty
	position = Position{2, 4}

	expected = "."
	actual = labMap.getCellAtPosition(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	// Test out of bounds
	position = Position{10, 10}

	expected = ""
	actual = labMap.getCellAtPosition(position)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestLabMapGetStartingPosition(t *testing.T) {
	inputs := getTestInputs()

	labMap := inputs.labMap

	expected := Position{4, 6}
	actual := labMap.getStartingPosition()

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestGuardianHasLooped(t *testing.T) {
	performTest := func(inputs ProblemInput, position Position, expected bool) {
		guardian := NewGuardian(inputs.labMap)
		inputs.labMap.area[position.y][position.x] = "#"
		guardian.visitTheLab()
		inputs.labMap.area[position.y][position.x] = "."
		actual := guardian.hasLooped()

		if actual != expected {
			t.Errorf("Expected %v but got %v . Values : %v", expected, actual, position)
		}
	}

	inputs := getTestInputs()

	performTest(inputs, Position{3, 6}, true)
	performTest(inputs, Position{6, 7}, true)
	performTest(inputs, Position{7, 7}, true)
	performTest(inputs, Position{1, 8}, true)
	performTest(inputs, Position{3, 8}, true)
	performTest(inputs, Position{7, 9}, true)

	performTest(inputs, Position{5, 4}, false)

}
