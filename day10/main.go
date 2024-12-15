package main

import (
	"fmt"
	"image"
)

type TopographicMap struct {
	area [][]int
}

type Trail struct {
	head           image.Point
	tails          []image.Point
	topographicMap TopographicMap
}

type ProblemInput struct {
	topographicMap TopographicMap
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return inputs.topographicMap.getScore()
}

func secondPart(inputs ProblemInput) int {
	return 0
}

func loadInputs(filename string) (inputs ProblemInput) {
	inputs = ProblemInput{}

	// todo

	return inputs
}

func getTopographicMapScore(topographicMap TopographicMap) int {
	total := 0

	trails := topographicMap.getTrails()
	for _, trail := range trails {
		trail.findTails()
		total += trail.getScore()
	}

	return total
}

// Trail methods

func NewTrail(head image.Point, topographicMap TopographicMap) Trail {
	return Trail{
		head:           head,
		tails:          []image.Point{},
		topographicMap: topographicMap,
	}
}

func (t *Trail) findTails() {
	tails := []image.Point{}

	isWithinBounds := func(position image.Point) bool {
		return position.X >= 0 && position.X < len(t.topographicMap.area[0]) && position.Y >= 0 && position.Y < len(t.topographicMap.area)
	}

	isHigherThanPosition := func(firstPosition, secondPosition image.Point) bool {
		if !isWithinBounds(firstPosition) || !isWithinBounds(secondPosition) {
			return false
		}

		positionValue := t.topographicMap.area[firstPosition.Y][firstPosition.X]
		secondPositionValue := t.topographicMap.area[secondPosition.Y][secondPosition.X]
		return secondPositionValue == positionValue+1
	}

	getTailPotentials := func(position image.Point) []image.Point {
		north := image.Point{position.X, position.Y - 1}
		south := image.Point{position.X, position.Y + 1}
		east := image.Point{position.X + 1, position.Y}
		west := image.Point{position.X - 1, position.Y}

		potentials := []image.Point{}

		// if they're not out of bounds and their value is 1 higher than the current position,
		// then they're potential tails

		if isHigherThanPosition(position, north) {
			potentials = append(potentials, north)
		}

		if isHigherThanPosition(position, south) {
			potentials = append(potentials, south)
		}

		if isHigherThanPosition(position, east) {
			potentials = append(potentials, east)
		}

		if isHigherThanPosition(position, west) {
			potentials = append(potentials, west)
		}

		return potentials
	}

	// get all potential tails for the head
	// if there are no potential tails, then don't keep the tails empty.
	// if there are potential tails, check if they have potential tails, and so on.
	// when there's no more potential tails, then we have a tail.

	tails = []image.Point{}
	for {
		potentials := getTailPotentials(t.head)
		if len(potentials) == 0 {
			break
		}

		// look for the next potential tail
		for _, potential := range potentials {
			nextPotentials := getTailPotentials(potential)

			for len(nextPotentials) > 0 {
				potentials = nextPotentials
				nextPotentials = getTailPotentials(potentials[0])
			}

			tails = append(tails, potential)
		}

	}

	t.tails = tails
}

func (t Trail) getScore() int {
	total := len(t.tails)

	return total
}

// TopographicMap methods

func (t TopographicMap) getTrails() []Trail {
	trails := []Trail{}

	// Loop through the topographic map and find all trail heads : 0
	for y, row := range t.area {
		for x, cell := range row {
			if cell != 0 {
				continue
			}

			position := image.Point{x, y}
			trail := NewTrail(position, t)
			trails = append(trails, trail)
		}
	}

	return trails
}

func (t TopographicMap) getScore() int {
	total := 0

	trails := t.getTrails()
	for _, trail := range trails {
		trail.findTails()
		total += trail.getScore()
	}

	return total
}
