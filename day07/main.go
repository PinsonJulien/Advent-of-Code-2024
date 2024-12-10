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
	Concatenate
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
	return getTotalCalibrationResultWithConcatenate(inputs)
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

func getTotalCalibrationResultWithConcatenate(input ProblemInput) int {
	total := 0
	for _, equation := range getValidEquationsWithConcatenate(input) {
		total = total + equation.result
	}

	return total
}

func getValidEquations(input ProblemInput) []Equation {
	validEquations := []Equation{}

	for _, equation := range input.equations {
		if equation.isPossible([]Operator{Addition, Multiplication}) {
			validEquations = append(validEquations, equation)
		}
	}

	return validEquations
}

func equationContains(equations []Equation, equation Equation) bool {
	for _, eq := range equations {
		if eq.Equals(equation) {
			return true
		}
	}

	return false
}

func getValidEquationsWithConcatenate(input ProblemInput) []Equation {
	validEquations := getValidEquations(input)

	// first, we check if the equation is possible with only addition and multiplication

	// then, we check if the equation is possible with addition, multiplication and concatenation
	invalidEquations := []Equation{}
	// Only add the equations that are not valid with only addition and multiplication
	for _, equation := range input.equations {
		if !equationContains(validEquations, equation) {
			invalidEquations = append(invalidEquations, equation)
		}
	}

	for _, equation := range invalidEquations {
		if equation.isPossible([]Operator{Addition, Multiplication, Concatenate}) {
			validEquations = append(validEquations, equation)
		}
	}

	return validEquations
}

// Equation methods

func (e *Equation) isPossible(allowedOperators []Operator) bool {
	calculateNumberOfPossibilities := func() int {
		size := len(allowedOperators)
		numberOfOperators := len(e.numbers) - 1
		return int(math.Pow(float64(size), float64(numberOfOperators)))
	}

	intToBaseString := func(n int64) string {
		base := len(allowedOperators)
		// convert the integer to a formatted string of base n
		return strconv.FormatInt(n, base)
	}

	determineOperators := func(n int) []Operator {
		operators := []Operator{}
		// convert the integer to a formatted string of base n
		baseStr := intToBaseString(int64(n))
		for _, c := range baseStr {
			operatorIndex, _ := strconv.Atoi(string(c))
			operator := allowedOperators[operatorIndex]
			operators = append(operators, operator)
		}

		// Add the missing operators before the firsts
		firstOperator := allowedOperators[0]
		for i := len(operators); i < len(e.numbers)-1; i++ {
			operators = append([]Operator{firstOperator}, operators...)
		}

		return operators
	}

	for i := 0; i < calculateNumberOfPossibilities(); i++ {
		operators := determineOperators(i)

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

func (e *Equation) Equals(equation Equation) bool {
	if e.result != equation.result {
		return false
	}

	if len(e.numbers) != len(equation.numbers) {
		return false
	}

	for i, number := range e.numbers {
		if number != equation.numbers[i] {
			return false
		}
	}

	return true
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
		case Concatenate:
			// Concatenate the value to the total
			totalStr := strconv.Itoa(total)
			valueStr := strconv.Itoa(value)
			totalStr = totalStr + valueStr
			total, _ = strconv.Atoi(totalStr)
		}
	}

	return total
}
