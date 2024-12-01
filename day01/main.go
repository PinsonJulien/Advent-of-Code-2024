package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	First  int
	Second int
}

func main() {
	// print inputs
	firstColumn, secondColumn := LoadInputs("inputs.txt")

	// print first part solution
	fmt.Println("First part solution: ", FirstPart(firstColumn, secondColumn))
	fmt.Println("Second part solution: ", SecondPart(firstColumn, secondColumn))
}

func FirstPart(firstColumn []int, secondColumn []int) int {
	pairs := GetPairs(firstColumn, secondColumn)
	return CalculateSumOfDistances(pairs)
}

func SecondPart(firstColumn []int, secondColumn []int) int {
	return CalculateSimilarityScore(firstColumn, secondColumn)
}

func LoadInputs(filename string) (firstColumn []int, secondColumn []int) {
	// reading inputs.txt as 2 arrays of integers, one for each column
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		// It should split "3   4\r" into ["3", "4"]
		// and then convert them to integers
		columns := strings.Fields(line)
		first, err := strconv.Atoi(columns[0])
		if err != nil {
			fmt.Println("Error converting to integer")
			os.Exit(1)
		}

		second, err := strconv.Atoi(columns[1])
		if err != nil {
			fmt.Println("Error converting to integer")
			os.Exit(1)
		}

		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)
	}

	return firstColumn, secondColumn
}

func GetPairs(firstColumn []int, secondColumn []int) []Pair {

	// Sort the two arrays lowest to highest int.
	slices.Sort(firstColumn)
	slices.Sort(secondColumn)

	// it should return an array of Pair
	var pairs []Pair
	for i := 0; i < len(firstColumn); i++ {
		pair := Pair{First: firstColumn[i], Second: secondColumn[i]}
		pairs = append(pairs, pair)
	}

	return pairs
}

func CalculateSumOfDistances(pairs []Pair) int {
	// returns the sum of all distances
	var sum int
	for _, pair := range pairs {
		sum += pair.Distance()
	}
	return sum
}

func CountOccurrences(value int, column []int) int {
	// returns the number of times a value appears in the column
	var count int
	for _, v := range column {
		if v == value {
			count++
		}
	}

	return count
}

func CalculateSimilarityScore(firstColumn []int, secondColumn []int) int {
	var score int
	for _, value := range firstColumn {
		score += value * CountOccurrences(value, secondColumn)
	}
	return score
}

/************ Pair methods ************/

func (p Pair) Distance() int {
	// Measure how far apart the two numbers are
	// Highest number - Lowest number
	if p.First > p.Second {
		return p.First - p.Second
	} else {
		return p.Second - p.First
	}
}
