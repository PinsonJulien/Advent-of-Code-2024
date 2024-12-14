package main

import (
	"fmt"
	"image"
	"maps"
	"os"
	"slices"
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
	return countAllAntiNodesWithResonantHarmonics(inputs)
}

func loadInputs(filename string) (inputs ProblemInput) {
	inputs = NewProblemInput()

	input, _ := os.ReadFile(filename)
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			point := image.Pt(x, y)
			inputs.addPoint(point)
			if r != '.' {
				inputs.addFrequency(r, point)
			}
		}
	}

	return inputs
}

func countAllAntiNodes(inputs ProblemInput) int {
	return len(getAllAntiNodes(inputs))
}

func getAllAntiNodes(inputs ProblemInput) []image.Point {
	antiNodes := map[image.Point]struct{}{}
	for _, antennas := range inputs.frequencies {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}
				a := a2.Add(a2.Sub(a1))
				if inputs.points[a] {
					antiNodes[a] = struct{}{}
				}
			}
		}
	}

	return slices.Collect(maps.Keys(antiNodes))
}

func countAllAntiNodesWithResonantHarmonics(inputs ProblemInput) int {
	return len(getAllAntiNodesWithResonantHarmonics(inputs))
}

func getAllAntiNodesWithResonantHarmonics(inputs ProblemInput) []image.Point {
	antiNodes := map[image.Point]struct{}{}
	for _, antennas := range inputs.frequencies {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}
				for d := a2.Sub(a1); inputs.points[a2]; a2 = a2.Add(d) {
					antiNodes[a2] = struct{}{}
				}
			}
		}
	}

	return slices.Collect(maps.Keys(antiNodes))
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
