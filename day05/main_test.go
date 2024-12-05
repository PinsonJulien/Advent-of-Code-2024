package main

import (
	"reflect"
	"testing"
)

const DefaultTestInputFile = "inputs.txt.example"

func getTestInputs() ProblemInput {
	/*
		47|53
		97|13
		97|61
		97|47
		75|29
		61|13
		75|53
		29|13
		97|29
		53|29
		61|53
		97|53
		61|29
		47|13
		75|47
		97|75
		47|61
		75|61
		47|29
		75|13
		53|13

		75,47,61,53,29
		97,61,53,29,13
		75,29,13
		75,97,47,61,53
		61,13,29
		97,13,75,29,47
	*/

	return ProblemInput{
		pageOrderingRules: []PageOrderingRule{
			{47, 53},
			{97, 13},
			{97, 61},
			{97, 47},
			{75, 29},
			{61, 13},
			{75, 53},
			{29, 13},
			{97, 29},
			{53, 29},
			{61, 53},
			{97, 53},
			{61, 29},
			{47, 13},
			{75, 47},
			{97, 75},
			{47, 61},
			{75, 61},
			{47, 29},
			{75, 13},
			{53, 13},
		},
		safetyManuals: []SafetyManual{
			{[]int{75, 47, 61, 53, 29}},
			{[]int{97, 61, 53, 29, 13}},
			{[]int{75, 29, 13}},
			{[]int{75, 97, 47, 61, 53}},
			{[]int{61, 13, 29}},
			{[]int{97, 13, 75, 29, 47}},
		},
	}
}

func TestLoadInputs(t *testing.T) {
	inputs := getTestInputs()
	actual := loadInputs(DefaultTestInputFile)
	if !reflect.DeepEqual(actual, inputs) {
		t.Errorf("Expected %v but got %v", inputs, actual)
	}
}

func TestSumAllMiddlePageOfValidSafetyManuals(t *testing.T) {
	inputs := getTestInputs()
	actual := sumAllMiddlePageOfValidSafetyManuals(inputs)
	expected := 61 + 53 + 29
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSumAllMiddlePageOfCorrectedSafetyManuals(t *testing.T) {
	inputs := getTestInputs()
	actual := sumAllMiddlePageOfCorrectedSafetyManuals(inputs)
	expected := 47 + 29 + 47
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSafetyManualGetMiddlePage(t *testing.T) {
	performTest := func(pages []int, expected int) {
		safetyManual := SafetyManual{pages}
		actual := safetyManual.getMiddlePage()
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
		}
	}

	performTest([]int{75, 47, 61, 53, 29}, 61)
	performTest([]int{97, 61, 53, 29, 13}, 53)
	performTest([]int{75, 29, 13}, 29)
	performTest([]int{75, 97, 47, 61, 53}, 47)
}

func TestSafetyManualIsValidAccordingToRules(t *testing.T) {
	testInputs := getTestInputs()

	performTest := func(pages []int, rules []PageOrderingRule, expected bool) {
		safetyManual := SafetyManual{pages}
		actual := safetyManual.isValidAccordingToRules(rules)
		if actual != expected {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}

	rules := testInputs.pageOrderingRules

	performTest(
		[]int{75, 47, 61, 53, 29}, rules,
		true,
	)

	performTest(
		[]int{97, 61, 53, 29, 13}, rules,
		true,
	)

	performTest(
		[]int{75, 29, 13}, rules,
		true,
	)

	performTest(
		[]int{75, 97, 47, 61, 53}, rules,
		false,
	)

	performTest(
		[]int{61, 13, 29}, rules,
		false,
	)

	performTest(
		[]int{97, 13, 75, 29, 47}, rules,
		false,
	)
}

func TestSafetyManualIsValidAccordingToRule(t *testing.T) {
	performTest := func(pages []int, rule PageOrderingRule, expected bool) {
		safetyManual := SafetyManual{pages}
		actual := safetyManual.isValidAccordingToRule(rule)
		if actual != expected {
			t.Errorf("Expected %v but got %v. SafetyManual: %v, Rule: %v", expected, actual, safetyManual, rule)
		}
	}

	pages := []int{75, 47, 61, 53, 29}

	// 75 is in correct order, because all other pages are after 75
	rule := PageOrderingRule{75, 47}
	performTest(pages, rule, true)
	rule = PageOrderingRule{75, 61}
	performTest(pages, rule, true)
	rule = PageOrderingRule{75, 53}
	performTest(pages, rule, true)
	rule = PageOrderingRule{75, 29}
	performTest(pages, rule, true)

	// 47 is in correct order, because 75 is before 47 and all other pages are after 47
	rule = PageOrderingRule{47, 61}
	performTest(pages, rule, true)
	rule = PageOrderingRule{47, 53}
	performTest(pages, rule, true)
	rule = PageOrderingRule{47, 29}
	performTest(pages, rule, true)

	// 61 is in correct order, because 75 and 47 are before 61 and all other pages are after 61
	rule = PageOrderingRule{61, 53}
	performTest(pages, rule, true)
	rule = PageOrderingRule{61, 29}
	performTest(pages, rule, true)

	// 53 is in correct order, because 75, 47 and 61 are before 53 and 29 is after 53
	rule = PageOrderingRule{53, 29}
	performTest(pages, rule, true)

	// 29 is in correct order, because 75, 47, 61 and 53 are before 29
	rule = PageOrderingRule{29, 13}
	performTest(pages, rule, true)

	pages = []int{75, 97, 47, 61, 53}

	// 97 is not in correct order, because 75 is before 97
	rule = PageOrderingRule{97, 75}
	performTest(pages, rule, false)

	pages = []int{61, 13, 29}
	// 13 is not in correct order, because 29 is after 13
	rule = PageOrderingRule{29, 13}
	performTest(pages, rule, false)

	// Correct order should be : 97, 75, 47, 29, 13
	pages = []int{97, 13, 75, 29, 47}
	// 97 is in correct order, because 97 is before every other page
	rule = PageOrderingRule{97, 13}
	performTest(pages, rule, true)

	// 13 is not in correct order, because 75 is before 13
	rule = PageOrderingRule{75, 13}
	performTest(pages, rule, false)

	// 29 is in correct order, because 97, 75, 47 are before 29 and 13 is after 29
	rule = PageOrderingRule{29, 13}
	performTest(pages, rule, false)

	// 47 is in correct order, because 97, 75 are before 47 and 29, 13 are after 47
	rule = PageOrderingRule{47, 29}
	performTest(pages, rule, false)
}

func TestSortByPageOrderingRules(t *testing.T) {
	performTest := func(pages []int, rules []PageOrderingRule, expected []int) {
		safetyManual := SafetyManual{pages}
		safetyManual.sortByPageOrderingRules(rules)
		actual := safetyManual.pages
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v but got %v. SafetyManual: %v, Rules: %v", expected, actual, safetyManual, rules)
		}
	}

	rules := getTestInputs().pageOrderingRules

	// Do not sort, because all pages are in correct order
	pages := []int{75, 47, 61, 53, 29}
	performTest(pages, rules, pages)

	// Do not sort, because all pages are in correct order
	pages = []int{97, 61, 53, 29, 13}
	performTest(pages, rules, pages)

	// Do not sort, because all pages are in correct order
	pages = []int{75, 29, 13}
	performTest(pages, rules, pages)

	// Sort, because 97 is before 75
	pages = []int{75, 97, 47, 61, 53}
	performTest(pages, rules, []int{97, 75, 47, 61, 53})

	// Sort, because 13 is before 29
	pages = []int{61, 13, 29}
	performTest(pages, rules, []int{61, 29, 13})

	// Sort, because 97 is before 75
	pages = []int{97, 13, 75, 29, 47}
	performTest(pages, rules, []int{97, 75, 47, 29, 13})
}
