package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	// North direction
	North Direction = iota
	// East direction
	East
	// South direction
	South
	// West direction
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
	return len(inputs.labMap.getPositionsWhichWouldCauseALoop())
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

func (LabMap *LabMap) getPositionsWhichWouldCauseALoop() []Position {
	originalGuardian := NewGuardian(*LabMap)
	originalGuardian.visitTheLab()

	positionsWhichWouldCauseALoop := map[Position]bool{}

	// Loop through the guardian's path and find the positions that would cause a loop

	// for each position visited by the guardian, add an obstacle and see if the guardian would visit it again
	for _, position := range originalGuardian.visitedPositions {
		// don't do anything if the position is the starting position
		if position == originalGuardian.LabMap.getStartingPosition() {
			continue
		}

		LabMap.addObstacle(position)
		guardian := NewGuardian(*LabMap)

		guardian.visitTheLab()

		// If the guardian visited the same position twice, it would cause a loop
		if guardian.isLooping {
			positionsWhichWouldCauseALoop[position] = true
		}

		// Remove the obstacle
		LabMap.removeObstacle(position)
	}

	positions := []Position{}
	for position := range positionsWhichWouldCauseALoop {
		positions = append(positions, position)
	}

	return positions
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

	// We need at least 6 positions to detect a meaningful loop
	if len(positions) < 6 {
		return false
	}

	// Start from the end and look for potential loop patterns
	for loopLength := 2; loopLength <= len(positions)/2; loopLength++ {
		// Check if the last two segments are identical
		if isIdenticalSegments(positions[len(positions)-loopLength*2:], loopLength) {
			return true
		}
	}

	return false
}

func isIdenticalSegments(positions []Position, segmentLength int) bool {
	// Split the slice into two equal segments
	firstSegment := positions[:segmentLength]
	secondSegment := positions[segmentLength : segmentLength*2]

	// Compare each position in the segments
	for i := 0; i < segmentLength; i++ {
		if firstSegment[i] != secondSegment[i] {
			return false
		}
	}

	return true
}
