package main

import "testing"

type TestEquivalence struct {
	ExpectedResult bool
	Expression1    string
	Expression2    string
}

func TestLogicalEquivalenceCalcualtor(t *testing.T) {
	tests := []TestEquivalence{}
	for _, test := range tests {
		result, _ := LogicalEquivalenceCalculator(test.Expression1, test.Expression2)
		if test.ExpectedResult != result {
			t.Errorf("Expected: %v\nGot: %v\n", test.ExpectedResult, result)
		}
	}
}
