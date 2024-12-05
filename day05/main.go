package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageOrderingRule struct {
	pageX int
	pageY int
}

type SafetyManual struct {
	pages []int
}

type ProblemInput struct {
	pageOrderingRules []PageOrderingRule
	safetyManuals     []SafetyManual
}

func main() {
	// print inputs
	inputs := loadInputs("inputs.txt")

	// print first part solution
	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return sumAllMiddlePageOfValidSafetyManuals(inputs)
}

func secondPart(inputs ProblemInput) int {
	return sumAllMiddlePageOfCorrectedSafetyManuals(inputs)
}

func loadInputs(filename string) (inputs ProblemInput) {
	// Parse data from file

	// First values are PageOrderingRules, 1 per line.
	// Format : %d|%d
	// Then, after an empty line
	// Second values are SafetyManuals, 1 per line.
	// Format : %d,%d,%d,%d,%d

	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")

	rowCount := len(lines)
	i := 0
	for ; i < rowCount; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		pageX, _ := strconv.Atoi(parts[0])
		pageY, _ := strconv.Atoi(parts[1])

		inputs.pageOrderingRules = append(inputs.pageOrderingRules, PageOrderingRule{pageX, pageY})
	}

	i++
	for ; i < rowCount; i++ {
		line := strings.TrimSpace(lines[i])
		parts := strings.Split(line, ",")
		safetyManual := SafetyManual{pages: []int{}}
		for _, part := range parts {
			page, _ := strconv.Atoi(part)
			safetyManual.pages = append(safetyManual.pages, page)
		}

		inputs.safetyManuals = append(inputs.safetyManuals, safetyManual)
	}

	return inputs
}

func sumAllMiddlePageOfValidSafetyManuals(inputs ProblemInput) int {
	total := 0

	for _, safetyManual := range inputs.safetyManuals {
		if safetyManual.isValidAccordingToRules(inputs.pageOrderingRules) {
			total += safetyManual.getMiddlePage()
		}
	}

	return total
}

func sumAllMiddlePageOfCorrectedSafetyManuals(inputs ProblemInput) int {
	total := 0

	for _, safetyManual := range inputs.safetyManuals {
		// We don't include valid safety manuals
		if safetyManual.isValidAccordingToRules(inputs.pageOrderingRules) {
			continue
		}

		// We correct the safety manual
		safetyManual.sortByPageOrderingRules(inputs.pageOrderingRules)
		total += safetyManual.getMiddlePage()
	}

	return total
}

// SafetyManual methods
func (s *SafetyManual) getMiddlePage() int {
	// Get the middle value of the pages, always odd.
	middleIndex := len(s.pages) / 2
	return s.pages[middleIndex]
}

func (s *SafetyManual) isValidAccordingToRules(pageOrderingRules []PageOrderingRule) bool {
	for _, rule := range pageOrderingRules {
		if !s.isValidAccordingToRule(rule) {
			return false
		}
	}

	return true
}

func (s *SafetyManual) isValidAccordingToRule(rule PageOrderingRule) bool {
	// Check if all pages are in the correct order
	// pageX is before pageY

	// Find the index of pageX and pageY
	pageXIndex := slices.Index(s.pages, rule.pageX)
	pageYIndex := slices.Index(s.pages, rule.pageY)

	// If pageX is not in the pages, it's always valid
	if pageXIndex == -1 {
		return true
	}

	// If pageY is not in the pages, it's always valid
	if pageYIndex == -1 {
		return true
	}

	// If pageX is before pageY, it's valid
	return pageXIndex < pageYIndex
}

func (s *SafetyManual) sortByPageOrderingRules(pageOrderingRules []PageOrderingRule) {
	// Sort the pages according to the rules

	for !s.isValidAccordingToRules(pageOrderingRules) {
		for _, rule := range pageOrderingRules {
			// Find the index of pageX and pageY
			pageXIndex := slices.Index(s.pages, rule.pageX)
			pageYIndex := slices.Index(s.pages, rule.pageY)

			// If pageX is not in the pages, it's always valid
			if pageXIndex == -1 {
				continue
			}

			// If pageY is not in the pages, it's always valid
			if pageYIndex == -1 {
				continue
			}

			// If pageX is before pageY, it's valid
			if pageXIndex < pageYIndex {
				continue
			}

			// Swap the values
			s.pages[pageXIndex], s.pages[pageYIndex] = s.pages[pageYIndex], s.pages[pageXIndex]
		}
	}
}
