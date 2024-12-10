package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	Addition Operator = iota
	Multiplication
)

type Equation struct {
	result  int
	numbers []int
}

type ProblemInput struct {
	equations []Equation
}

type Calculator struct {
	values    []int
	operators []Operator
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return getTotalCalibrationResult(inputs)
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

	equations := []Equation{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimSpace(line)

		// Parse the line like so :
		// 161011: 16 10 13
		// result: 161011
		// numbers: [16, 10, 13]

		// Split the line into two parts
		parts := strings.Split(line, ":")
		result, _ := strconv.Atoi(parts[0])

		strNumbers := strings.Split(parts[1], " ")
		numbers := []int{}
		for _, strNumber := range strNumbers {
			strNumber = strings.TrimSpace(strNumber)
			number, _ := strconv.Atoi(strNumber)
			numbers = append(numbers, number)
		}

		equation := Equation{result: result, numbers: numbers}
		equations = append(equations, equation)
	}

	inputs = ProblemInput{equations: equations}

	return inputs
}

func getTotalCalibrationResult(input ProblemInput) int {
	total := 0
	for _, equation := range getValidEquations(input) {
		total = total + equation.result
	}

	return total
}

func getValidEquations(input ProblemInput) []Equation {
	validEquations := []Equation{}

	for _, equation := range input.equations {
		if equation.isPossible() {
			validEquations = append(validEquations, equation)
		}
	}

	return validEquations
}

func convertIntegerToBinaryString(value int) string {
	n := int64(value)
	return strconv.FormatInt(n, 2)
}

func getOperatorsByBinaryString(binaryStr string, expectedLength int) []Operator {
	operators := []Operator{}

	for _, c := range binaryStr {
		operator := Addition
		if c != '0' {
			operator = Multiplication
		}

		operators = append(operators, operator)
	}

	// Add the missing operators before the firsts
	for i := len(operators); i < expectedLength; i++ {
		operators = append([]Operator{Addition}, operators...)
	}

	return operators
}

// Equation methods

func (e *Equation) isPossible() bool {
	// Check if the equation is possible to calculate
	expectedNumberOfOperators := len(e.numbers) - 1
	maxIterations := int(math.Pow(float64(2), float64(expectedNumberOfOperators)))

	// use binary to determine the operators for each loops
	for i := 0; i < maxIterations; i++ {
		binaryStr := convertIntegerToBinaryString(i)
		operators := getOperatorsByBinaryString(binaryStr, expectedNumberOfOperators)

		calculator := NewCalculator(
			e.numbers,
			operators,
		)

		equationResult := calculator.calculate()

		if equationResult == e.result {
			return true
		}
	}

	return false
}

// Calculator methods

func NewCalculator(values []int, operators []Operator) Calculator {
	return Calculator{
		values:    values,
		operators: operators,
	}
}

func (c *Calculator) calculate() int {
	// calculate left to right the positions with specified operator.
	total := c.values[0]
	valuesWithoutFirst := c.values[1:]

	for i, operator := range c.operators {
		value := valuesWithoutFirst[i]

		switch operator {
		case Addition:
			total = total + value
		case Multiplication:
			total = total * value
		}
	}

	return total
}
