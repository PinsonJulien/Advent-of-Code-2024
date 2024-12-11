package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type ProblemInput struct {
	points      map[image.Point]bool
	frequencies map[rune][]image.Point
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return countAllAntiNodes(inputs)
}

func secondPart(inputs ProblemInput) int {
	return 0
}

func loadInputs(filename string) (inputs ProblemInput) {
	inputs = NewProblemInput()

	input, _ := os.ReadFile(filename)

	// Parse the data here
	fileContent := string(input)
	for y, s := range strings.Fields(fileContent) {
		for x, r := range s {
			inputs.addPoint(image.Pt(x, y))
			if r != '.' {
				inputs.addFrequency(r, image.Pt(x, y))
			}
		}
	}

	return inputs
}

func countAllAntiNodes(inputs ProblemInput) int {
	return len(getAllAntiNodes(inputs))
}

func getAllAntiNodes(inputs ProblemInput) []image.Point {
	antiNodes := []image.Point{}

	for _, antennas := range inputs.frequencies {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}

				antiNode := a2.Add(a2.Sub(a1))
				if inputs.points[antiNode] {
					antiNodes = append(antiNodes, antiNode)
				}
			}
		}
	}

	return antiNodes
}

// ProblemInput methods

func NewProblemInput() ProblemInput {
	return ProblemInput{
		points:      map[image.Point]bool{},
		frequencies: map[rune][]image.Point{},
	}
}

func (p *ProblemInput) addPoint(point image.Point) {
	p.points[point] = true
}

func (p *ProblemInput) addFrequency(frequency rune, point image.Point) {
	p.frequencies[frequency] = append(p.frequencies[frequency], point)
}
