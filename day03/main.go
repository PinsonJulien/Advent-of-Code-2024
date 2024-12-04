package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Pair class
type Pair struct {
	first   int
	second  int
	enabled bool
}

func main() {
	pairs := loadInputs("inputs.txt")

	// print first part solution
	fmt.Println("First part solution: ", firstPart(pairs))
	fmt.Println("Second part solution: ", secondPart(pairs))
}

func firstPart(pairs []Pair) int {
	return multiplyAndSumPairs(pairs, false)
}

func secondPart(pairs []Pair) int {
	return multiplyAndSumPairs(pairs, true)
}

func loadInputs(filename string) (pairs []Pair) {
	// It should return an array of Pair type.
	// Pair is based on mul(3, 4) values, so first column is 3 and second column is 4.
	// It should be able to prepare the array of Pair type by being able to extract "mul(val, val)"
	// from : xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
	// in this case, the expected pairs are :
	// Pair(2,4, true) ; Pair(5,5, false) ; Pair(11,8, false) ; Pair(8,5, true)

	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	fileContent := string(data)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	parts := re.FindAllString(fileContent, -1)
	// Always start with enabled
	enabled := true
	for _, part := range parts {
		// When we see do() or don't(), we change the enabled flag
		if part == "do()" {
			enabled = true
			continue
		} else if part == "don't()" {
			enabled = false
			continue
		}

		// Add mul() pairs
		re = regexp.MustCompile(`\d+`)
		matches := re.FindAllString(part, -1)

		first, err := strconv.Atoi(matches[0])
		if err != nil {
			continue
		}

		second, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}

		pairs = append(
			pairs,
			Pair{first, second, enabled},
		)
	}

	return pairs
}

func multiplyAndSumPairs(pairs []Pair, excludeDisabled bool) int {
	sum := 0
	for _, pair := range pairs {
		if excludeDisabled && !pair.enabled {
			continue
		}

		sum += pair.multiply()
	}
	return sum
}

// Pair methods
func (p Pair) multiply() int {
	return p.first * p.second
}
