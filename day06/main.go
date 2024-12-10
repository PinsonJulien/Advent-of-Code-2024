package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	x int
	y int
}

type LabMap struct {
	area [][]string
}

type ProblemInput struct {
	labMap LabMap
}

type Guardian struct {
	currentPosition  Position
	direction        Direction
	LabMap           LabMap
	leftTheLab       bool
	visitedPositions []Position
	isLooping        bool
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	guardian := NewGuardian(inputs.labMap)
	guardian.visitTheLab()

	return guardian.getUniquePositionsVisited()
}

func secondPart(inputs ProblemInput) int {
	// Create a copy of the lab map to avoid modifying the original
	loopPositions := inputs.labMap.getPositionsWhichWouldCauseALoop()
	return len(loopPositions)
}

func loadInputs(filename string) (inputs ProblemInput) {
	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")
	LabMapArea := [][]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimSpace(line)
		LabMapArea = append(LabMapArea, strings.Split(line, ""))
	}

	inputs.labMap = LabMap{area: LabMapArea}

	return inputs
}

// LabMap methods
func (LabMap *LabMap) getStartingPosition() Position {
	for y, row := range LabMap.area {
		for x, cell := range row {
			if cell == "^" {
				return Position{x, y}
			}
		}
	}

	panic("No starting position found")
}

func (LabMap *LabMap) getCellAtPosition(position Position) string {
	if LabMap.isPositionOutOfBounds(position) {
		return ""
	}

	return LabMap.area[position.y][position.x]
}

func (LabMap *LabMap) isPositionOutOfBounds(position Position) bool {
	return position.x < 0 || position.x >= len(LabMap.area[0]) || position.y < 0 || position.y >= len(LabMap.area)
}

func (LabMap *LabMap) isPositionWall(position Position) bool {
	return LabMap.getCellAtPosition(position) == "#"
}

func (labMap *LabMap) getPositionsWhichWouldCauseALoop() []Position {
	// Create a deep copy of the lab map to avoid modifying the original
	copiedLabMap := LabMap{area: make([][]string, len(labMap.area))}
	for i, row := range labMap.area {
		copiedLabMap.area[i] = make([]string, len(row))
		copy(copiedLabMap.area[i], row)
	}

	originalGuardian := NewGuardian(copiedLabMap)
	originalGuardian.visitTheLab()

	positionsWhichWouldCauseALoop := []Position{}

	// Loop through the positions that could potentially cause a loop
	fmt.Println("Length of visited positions", len(originalGuardian.visitedPositions))
	for i, position := range originalGuardian.visitedPositions {
		fmt.Println("Checking position", i, "of", len(originalGuardian.visitedPositions))
		// Skip the starting position
		if position == originalGuardian.LabMap.getStartingPosition() {
			continue
		}

		// Create a new copy of the lab map for each test
		testLabMap := LabMap{area: make([][]string, len(labMap.area))}
		for i, row := range labMap.area {
			testLabMap.area[i] = make([]string, len(row))
			copy(testLabMap.area[i], row)
		}

		// Add an obstacle
		testLabMap.area[position.y][position.x] = "#"

		// Test if this obstacle causes a loop
		guardian := NewGuardian(testLabMap)
		guardian.visitTheLab()

		// If the guardian is looping, add this position
		if guardian.isLooping {
			positionsWhichWouldCauseALoop = append(positionsWhichWouldCauseALoop, position)
		}
	}

	// return all unique positions

	postionsMap := make(map[Position]bool)
	for _, position := range positionsWhichWouldCauseALoop {
		postionsMap[position] = true

	}

	uniquePositions := []Position{}
	for position := range postionsMap {
		uniquePositions = append(uniquePositions, position)
	}

	return uniquePositions
}

func (LabMap *LabMap) addObstacle(position Position) {
	LabMap.area[position.y][position.x] = "#"
}

func (LabMap *LabMap) removeObstacle(position Position) {
	LabMap.area[position.y][position.x] = "."
}

// Guardian methods

func NewGuardian(labMap LabMap) Guardian {
	startPosition := labMap.getStartingPosition()

	return Guardian{
		currentPosition:  startPosition,
		direction:        North,
		LabMap:           labMap,
		leftTheLab:       false,
		visitedPositions: []Position{startPosition},
		isLooping:        false,
	}
}

func (guardian *Guardian) visitTheLab() {
	for !guardian.leftTheLab {
		guardian.moveForward()

		if guardian.hasLooped() {
			guardian.isLooping = true
			break
		}
	}
}

func (guardian *Guardian) moveForward() {
	if guardian.leftTheLab {
		return
	}

	newPosition := guardian.getNextPosition()

	if guardian.LabMap.isPositionOutOfBounds(newPosition) {
		guardian.leftTheLab = true
		return
	}

	if guardian.LabMap.isPositionWall(newPosition) {
		guardian.turn()
		newPosition = guardian.getNextPosition()
	}

	guardian.currentPosition = newPosition
	guardian.visitedPositions = append(guardian.visitedPositions, newPosition)
}

func (guardian *Guardian) getNextPosition() Position {
	newPosition := Position{
		x: guardian.currentPosition.x,
		y: guardian.currentPosition.y,
	}

	switch guardian.direction {
	case North:
		newPosition.y--
	case East:
		newPosition.x++
	case South:
		newPosition.y++
	case West:
		newPosition.x--
	}

	return newPosition
}

func (guardian *Guardian) turn() {
	switch guardian.direction {
	case North:
		guardian.direction = East
	case East:
		guardian.direction = South
	case South:
		guardian.direction = West
	case West:
		guardian.direction = North
	}
}

func (guardian *Guardian) getUniquePositionsVisited() int {
	uniquePositions := make(map[Position]bool)

	for _, position := range guardian.visitedPositions {
		uniquePositions[position] = true
	}

	return len(uniquePositions)
}

func (guardian *Guardian) hasLooped() bool {
	positions := guardian.visitedPositions
	countOfUniquePositions := guardian.getUniquePositionsVisited()

	// if the guardian has visited twice more than the length of unique positions, it has looped
	if len(positions) > countOfUniquePositions*2 {
		return true
	}

	return false
}

func isIdenticalSegments(firstSegment, secondSegment []Position) bool {
	// Ensure segments are of equal length
	if len(firstSegment) != len(secondSegment) {
		return false
	}

	// Compare each position in the segments
	for i := 0; i < len(firstSegment); i++ {
		if firstSegment[i] != secondSegment[i] {
			return false
		}
	}

	return true
}
