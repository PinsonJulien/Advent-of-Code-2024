package main

import (
	"fmt"
	"os"
	"strings"
)

// Direction enum
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
	currentPosition Position
	direction       Direction
	LabMap          LabMap
	leftTheLab      bool
	positionSet     map[Position]bool
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
	// for each row and column, check if the guardian would loop
	positions := []Position{}

	for y, row := range LabMap.area {
		for x, cell := range row {
			if cell != "." {
				continue
			}
			position := Position{x, y}
			if LabMap.wouldCauseALoop(position) {
				positions = append(positions, position)
			}
		}
	}

	return positions
}

func (LabMap *LabMap) wouldCauseALoop(position Position) bool {
	// Check if the position has a wall on north.
	// If it doesn't, then check if the position has a wall on south left

	northPosition := LabMap.getNextBlockPosition(position, North)
	if LabMap.isPositionWall(northPosition) {
		// Check if the position has a wall on west.
		underNorthPosition := Position{
			x: northPosition.x,
			y: northPosition.y + 1,
		}
		westPosition := LabMap.getNextBlockPosition(underNorthPosition, West)
		if !LabMap.isPositionWall(westPosition) {
			fmt.Println("westPosition", westPosition)
			return false
		}

		beforeWestPosition := Position{
			x: westPosition.x - 1,
			y: westPosition.y,
		}

		// Check if the position has a wall on south.
		southPosition := LabMap.getNextBlockPosition(beforeWestPosition, South)
		if !LabMap.isPositionWall(southPosition) {
			return false
		}

		return true
	}

	// Check if the position has a wall on south.
	leftPosition := Position{
		x: position.x - 1,
		y: position.y,
	}

	southPosition := LabMap.getNextBlockPosition(leftPosition, South)
	if LabMap.isPositionWall(southPosition) {
		// Check if the position has a wall on east.
		rightSouthPosition := Position{
			x: southPosition.x + 1,
			y: southPosition.y,
		}
		eastPosition := LabMap.getNextBlockPosition(rightSouthPosition, East)
		if !LabMap.isPositionWall(eastPosition) {
			return false
		}

		beforeEastPosition := Position{
			x: eastPosition.x - 1,
			y: eastPosition.y,
		}

		// Check if the position has a wall on north.
		northPosition := LabMap.getNextBlockPosition(beforeEastPosition, North)
		if !LabMap.isPositionWall(northPosition) {
			return false
		}

		return true
	}

	return false
}

func (LabMap *LabMap) getNextBlockPosition(currentPosition Position, direction Direction) Position {
	// Loop through the map until we find a wall in the given direction
	newPosition := Position{
		x: currentPosition.x,
		y: currentPosition.y,
	}

	// Move in the given direction until we find a wall
	for {
		switch direction {
		case North:
			newPosition.y--
		case East:
			newPosition.x++
		case South:
			newPosition.y++
		case West:
			newPosition.x--
		}

		if LabMap.isPositionOutOfBounds(newPosition) {
			break
		}

		if LabMap.isPositionWall(newPosition) {
			break
		}
	}

	return newPosition
}

// Guardian methods

func NewGuardian(labMap LabMap) Guardian {
	startPosition := labMap.getStartingPosition()
	positionSet := make(map[Position]bool)
	positionSet[startPosition] = true

	return Guardian{
		currentPosition: startPosition,
		direction:       North,
		LabMap:          labMap,
		leftTheLab:      false,
		positionSet:     positionSet,
	}
}

func (guardian *Guardian) visitTheLab() {
	previousPosition := guardian.currentPosition
	for !guardian.leftTheLab {
		guardian.moveForward()

		// If the guardian is stuck in a loop, stop.
		if previousPosition == guardian.currentPosition {
			break
		}

		previousPosition = guardian.currentPosition
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

	guardian.addPositionToSet()
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

func (guardian *Guardian) addPositionToSet() {
	if guardian.positionSet == nil {
		guardian.positionSet = make(map[Position]bool)
	}

	guardian.positionSet[guardian.currentPosition] = true
}

func (guardian *Guardian) getUniquePositionsVisited() int {
	return len(guardian.positionSet)
}
