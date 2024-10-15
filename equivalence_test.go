package main

import (
	"testing"
)

type TestEquivalence struct {
	ExpectedResult bool
	Expression1    string
	Expression2    string
}

func TestLogicalEquivalenceCalculator(t *testing.T) {
	tests := []TestEquivalence{
		{
			ExpectedResult: true,
			Expression1:    "x + xyz",
			Expression2:    "x",
		}, {
			ExpectedResult: false,
			Expression1:    "x + x'yz",
			Expression2:    "x",
		}, {
			ExpectedResult: true,
			Expression1:    "x'y'z + x'yz + xy'z + xyz",
			Expression2:    "z",
		}, {
			ExpectedResult: true,
			Expression1:    "x'y'z' + x'y'z + x'yz + xyz'",
			Expression2:    "x'y'+x'z+xyz'",
		}, {
			ExpectedResult: true,
			Expression1:    "(x+y'+z)(x'+y+z)(x'+y+z')(x'+y'+z')",
			Expression2:    "xyz'+x'y'+x'z",
		}, {
			ExpectedResult: false,
			Expression1:    "x'y'z' + x'y'z + x'yz + xyz'",
			Expression2:    "x'+x'z+xyz'",
		}, {
			ExpectedResult: true,
			Expression1:    "y",
			Expression2:    "x'yz'+x'yz+xyz'+xyz",
		}, {
			ExpectedResult: false,
			Expression1:    "x'yz'+x'yz+xyz'+xyz",
			Expression2:    "y'",
		},
	}
	for _, test := range tests {
		result, _ := LogicalEquivalenceCalculator(test.Expression1, test.Expression2, NoPrint)
		if test.ExpectedResult != result {
			t.Errorf(
				"\nExpression1: %v\nExpression2: %v\nExpected: %v\nGot: %v\n",
				test.Expression1,
				test.Expression2,
				test.ExpectedResult,
				result,
			)
		}
	}
}
