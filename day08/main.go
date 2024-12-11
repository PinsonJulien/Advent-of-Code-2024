package main

import (
	"fmt"
	"os"
	"strings"
)

type ProblemInput struct {
	cityMap [][]string
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return 0
}

func secondPart(inputs ProblemInput) int {
	return 0
}

func loadInputs(filename string) (inputs ProblemInput) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse the data here
	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")

	cityMap := [][]string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		cityMap = append(cityMap, strings.Split(line, ""))
	}

	inputs = ProblemInput{cityMap: cityMap}

	return inputs
}
