package main

import "testing"

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	equations := []Equation{
		{
			result:  190,
			numbers: []int{10, 19},
		},
		{
			result:  3267,
			numbers: []int{81, 40, 27},
		},
		{
			result:  83,
			numbers: []int{17, 5},
		},
		{
			result:  156,
			numbers: []int{15, 6},
		},
		{
			result:  7290,
			numbers: []int{6, 8, 6, 15},
		},
		{
			result:  161011,
			numbers: []int{16, 10, 13},
		},
		{
			result:  192,
			numbers: []int{17, 8, 14},
		},
		{
			result:  21037,
			numbers: []int{9, 7, 18, 13},
		},
		{
			result:  292,
			numbers: []int{11, 6, 16, 20},
		},
	}

	inputs := ProblemInput{equations: equations}

	return inputs
}

func TestLoadInputs(t *testing.T) {
	inputs := loadInputs(DefaultTestInputFile)

	if len(inputs.equations) != 9 {
		t.Errorf("Expected 9 equations, got %d", len(inputs.equations))
	}
}

func TestGetTotalCalibrationResult(t *testing.T) {
	inputs := getTestInputs()

	expected := 3749
	result := getTotalCalibrationResult(inputs)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetTotalCalibrationResultWithConcatenate(t *testing.T) {
	inputs := getTestInputs()

	expected := 11387
	result := getTotalCalibrationResultWithConcatenate(inputs)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetValidEquations(t *testing.T) {
	inputs := getTestInputs()

	expectedResults := []Equation{
		{
			result:  190,
			numbers: []int{10, 19},
		},
		{
			result:  3267,
			numbers: []int{81, 40, 27},
		},
		{
			result:  292,
			numbers: []int{11, 6, 16, 20},
		},
	}

	results := getValidEquations(inputs)

	expectedSize := len(expectedResults)
	actualSize := len(results)
	if actualSize != expectedSize {
		t.Errorf("Expected %d equations, got %d", expectedSize, actualSize)
	}

	for i := 0; i < expectedSize; i++ {
		expectedResult := expectedResults[i]
		actualResult := results[i]
		if expectedResult.result != actualResult.result {
			t.Errorf("Expected %d, got %d", expectedResult.result, actualResult.result)
		}
	}
}

func TestGetValidEquationsWithConcatenate(t *testing.T) {
	inputs := getTestInputs()

	expectedResults := []Equation{
		{
			result:  190,
			numbers: []int{10, 19},
		},
		{
			result:  3267,
			numbers: []int{81, 40, 27},
		},
		{
			result:  292,
			numbers: []int{11, 6, 16, 20},
		},
		{
			result:  156,
			numbers: []int{15, 6},
		},
		{
			result:  7290,
			numbers: []int{6, 8, 6, 15},
		},
		{
			result:  192,
			numbers: []int{17, 8, 14},
		},
	}

	results := getValidEquationsWithConcatenate(inputs)

	expectedSize := len(expectedResults)
	actualSize := len(results)
	if actualSize != expectedSize {
		t.Errorf("Expected %d equations, got %d", expectedSize, actualSize)
	}

	for i := 0; i < expectedSize; i++ {
		expectedResult := expectedResults[i]
		actualResult := results[i]
		if expectedResult.result != actualResult.result {
			t.Errorf("Expected %d, got %d", expectedResult.result, actualResult.result)
		}
	}
}

func TestEquationIsPossible(t *testing.T) {
	performTest := func(equation Equation, expected bool) {
		if equation.isPossible([]Operator{Addition, Multiplication}) != expected {
			message := "Equation should "
			if !expected {
				message += "not "
			}
			message += "be possible. Equation: %v"
			t.Errorf(message, equation)
		}
	}

	equation := Equation{
		result:  190,
		numbers: []int{10, 19},
	}
	expected := true
	performTest(equation, expected)

	equation = Equation{
		result:  3267,
		numbers: []int{81, 40, 27},
	}
	expected = true
	performTest(equation, expected)

	equation = Equation{
		result:  83,
		numbers: []int{17, 5},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  156,
		numbers: []int{15, 6},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  7290,
		numbers: []int{6, 8, 6, 15},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  161011,
		numbers: []int{16, 10, 13},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  192,
		numbers: []int{17, 8, 14},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  21037,
		numbers: []int{9, 7, 18, 13},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  292,
		numbers: []int{11, 6, 16, 20},
	}
	expected = true
	performTest(equation, expected)
}

func TestEquationIsPossibleWithConcatenate(t *testing.T) {
	performTest := func(equation Equation, expected bool) {
		if equation.isPossible([]Operator{Addition, Multiplication, Concatenate}) != expected {
			message := "Equation should "
			if !expected {
				message += "not "
			}
			message += "be possible. Equation: %v"
			t.Errorf(message, equation)
		}
	}

	equation := Equation{
		result:  190,
		numbers: []int{10, 19},
	}
	expected := true
	performTest(equation, expected)

	equation = Equation{
		result:  3267,
		numbers: []int{81, 40, 27},
	}
	expected = true
	performTest(equation, expected)

	equation = Equation{
		result:  83,
		numbers: []int{17, 5},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  156,
		numbers: []int{15, 6},
	}
	expected = true
	performTest(equation, expected)

	equation = Equation{
		result:  7290,
		numbers: []int{6, 8, 6, 15},
	}
	expected = true
	performTest(equation, expected)

	equation = Equation{
		result:  161011,
		numbers: []int{16, 10, 13},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  192,
		numbers: []int{17, 8, 14},
	}
	expected = true
	performTest(equation, expected)

	equation = Equation{
		result:  21037,
		numbers: []int{9, 7, 18, 13},
	}
	expected = false
	performTest(equation, expected)

	equation = Equation{
		result:  292,
		numbers: []int{11, 6, 16, 20},
	}
	expected = true
	performTest(equation, expected)
}
