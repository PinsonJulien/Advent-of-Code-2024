package main

import (
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestPairs() []Pair {
	return []Pair{
		{2, 4, true},
		{5, 5, false},
		{11, 8, false},
		{8, 5, true},
	}
}

func TestLoadInputs(t *testing.T) {

	// It should return an array of Pair type.
	// Pair is based on mul(3, 4) values, so first column is 3 and second column is 4.
	// It should be able to prepare the array of Pair type by being able to extract "mul(val, val)"
	// from : xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
	// in this case, the expected pairs are :
	// Pair(2,4) ; Pair(5,5) ; Pair(11,8) ; Pair(8,5)

	// However, the pairs should be enabled only if it is preceded by a "do()", if it is "don't()" then it should be disabled.
	// Therefore, the expected pairs are :
	// Pair(2,4, true) ; Pair(5,5, false) ; Pair(11,8, false) ; Pair(8,5, true)

	var expectedPairs = getTestPairs()

	pairs := loadInputs(DefaultTestInputFile)

	for i := 0; i < len(expectedPairs); i++ {
		if pairs[i].first != expectedPairs[i].first {
			t.Errorf("Expected pairs[%d].First to be %d, but got %d", i, expectedPairs[i].first, pairs[i].first)
		}
		if pairs[i].second != expectedPairs[i].second {
			t.Errorf("Expected pairs[%d].Second to be %d, but got %d", i, expectedPairs[i].second, pairs[i].second)
		}

		if pairs[i].enabled != expectedPairs[i].enabled {
			t.Errorf("Expected pairs[%d].Enabled to be %t, but got %t", i, expectedPairs[i].enabled, pairs[i].enabled)
		}
	}
}

func TestMultiplyAndSumPairs(t *testing.T) {
	// it should return 161 for the test pairs
	// 2*4 + 5*5 + 11*8 + 8*5 = 161
	expectedResult := 161
	pairs := getTestPairs()

	result := multiplyAndSumPairs(pairs, false)

	if result != expectedResult {
		t.Errorf("Expected result to be %d, but got %d", expectedResult, result)
	}
}

func TestMultiplyAndSumPairs_WithEnabledPairs(t *testing.T) {
	// it should return 48 for the test pairs, when we only want to include enabled pairs
	// 2*4 + 8*5 = 48
	expectedResult := 48

	pairs := getTestPairs()

	result := multiplyAndSumPairs(pairs, true)

	if result != expectedResult {
		t.Errorf("Expected result to be %d, but got %d", expectedResult, result)
	}
}

func TestPair_Multiply(t *testing.T) {
	// it should return 8, 25, 88, 40
	var ExpectedResults = []int{8, 25, 88, 40}

	pairs := getTestPairs()

	for i := 0; i < len(ExpectedResults); i++ {
		value := pairs[i].multiply()
		if pairs[i].multiply() != ExpectedResults[i] {
			t.Errorf("Expected pairs[%d].Multiply() to be %d, but got %d", i, ExpectedResults[i], value)
		}
	}
}
