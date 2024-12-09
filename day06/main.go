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
			fmt.Println("position", position)
			if LabMap.wouldCauseALoop(position) {
				positions = append(positions, position)
			}
		}
	}

	return positions
}

func (LabMap *LabMap) wouldCauseALoop(position Position) bool {
	// if position 3, 6, place a rock

	// temporarily add a rock at position.
	LabMap.area[position.y][position.x] = "#"
	defer func() {
		fmt.Println("changing back to .")
		LabMap.area[position.y][position.x] = "."
	}()

	// Get all the rocks in all directions
	rocks := LabMap.getRocksInAllDirections(position)

	// if there's less than 2 rocks, return false
	if len(rocks) < 2 {
		return false
	}

	// Check if position is one of the values of each all directions from all angles
	for _, rock := range rocks {
		walls := LabMap.getRocksInAllDirections(rock)
		if len(walls) < 2 {
			continue
		}

		fmt.Println("walls", walls)
		hasOriginalPositionAsWall := false
		hasAtLeastOneWall := false

		// Check if the position is one of the values of each all directions from all angles
		for _, wall := range walls {
			if wall == position {
				hasOriginalPositionAsWall = true
			} else {
				for _, r := range rocks {
					fmt.Println("has one ??")
					fmt.Println(r, wall)
					if wall == r {
						hasAtLeastOneWall = true
					}
				}
			}
		}

		if hasOriginalPositionAsWall && hasAtLeastOneWall {
			return true
		}
	}
	fmt.Println("done checking")
	return false

	// Check if the position has a wall on north.
	// If it doesn't, then check if the position has a wall on south left

	/*rightOfPosition := Position{
		x: position.x + 1,
		y: position.y,
	}
	northPosition := LabMap.getNextBlockPosition(rightOfPosition, North)
	if LabMap.isPositionWall(northPosition) {
		// Check if the position has a wall on west.
		underNorthPosition := Position{
			x: northPosition.x,
			y: northPosition.y + 1,
		}
		eastPosition := LabMap.getNextBlockPosition(underNorthPosition, East)
		if !LabMap.isPositionWall(eastPosition) {
			return false
		}

		beforeEastPosition := Position{
			x: eastPosition.x - 1,
			y: eastPosition.y,
		}

		// Check if the position has a wall on south.
		southPosition := LabMap.getNextBlockPosition(beforeEastPosition, South)
		if !LabMap.isPositionWall(southPosition) {
			return false
		}

		upSouthPosition := Position{
			x: southPosition.x,
			y: southPosition.y - 1,
		}

		// Check if the position has the same y position as the starting position
		if upSouthPosition.y == position.y {
			return true
		}

		return false
	}

	// Check if the position has a wall on south.
	leftPosition := Position{
		x: position.x - 1,
		y: position.y,
	}

	southPosition := LabMap.getNextBlockPosition(leftPosition, South)
	if LabMap.isPositionWall(southPosition) {
		// Check if the position has a wall on west.
		upSouthPosition := Position{
			x: southPosition.x,
			y: southPosition.y - 1,
		}
		eastPosition := LabMap.getNextBlockPosition(upSouthPosition, East)
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

		downNorthPosition := Position{
			x: northPosition.x,
			y: northPosition.y - 1,
		}

		// Check if the position has the same y position as the starting position
		if downNorthPosition.y == position.y {
			return true
		}

		return false
	}*/

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

func (LabMap *LabMap) getRocksInAllDirections(position Position) map[Direction]Position {
	rocks := make(map[Direction]Position)

	northPosition := LabMap.getRockInDirection(position, North)
	if LabMap.isPositionWall(northPosition) {
		rocks[North] = northPosition
	}

	eastPosition := LabMap.getRockInDirection(position, East)
	if LabMap.isPositionWall(eastPosition) {
		rocks[East] = eastPosition
	}

	southPosition := LabMap.getRockInDirection(position, South)
	if LabMap.isPositionWall(southPosition) {
		rocks[South] = southPosition
	}

	westPosition := LabMap.getRockInDirection(position, West)
	if LabMap.isPositionWall(westPosition) {
		rocks[West] = westPosition
	}

	return rocks
}

func (LabMap *LabMap) getRockInDirection(position Position, direction Direction) Position {
	newPosition := Position{
		x: position.x,
		y: position.y,
	}

	switch direction {
	case North:
		newPosition.x++
	case East:
		newPosition.y++
	case South:
		newPosition.x--
	case West:
		newPosition.y--
	}

	return LabMap.getNextBlockPosition(newPosition, direction)
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
