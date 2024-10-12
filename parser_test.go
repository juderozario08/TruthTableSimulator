package main

import (
	"slices"
	"testing"
)

type ParseTermsTest struct {
	ExpectedTerms []Term
	ExpectedIsPOS bool
	Question      string
}

func ParserTest(t *testing.T) {
	tests := []ParseTermsTest{}
	for _, test := range tests {
		tokens, _, _ := GenerateTokensAndStates(test.Question)
		terms, isPOS := ParseTerms(&tokens)
		if !termEqual(terms, test.ExpectedTerms) {
			t.Errorf("Expected: %v\nGot: %v", test.ExpectedTerms, terms)
		}
		if isPOS == test.ExpectedIsPOS {
			t.Errorf("Expected: %v\nGot: %v", test.ExpectedIsPOS, isPOS)
		}
	}
}

func termEqual(t1 []Term, t2 []Term) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if !slices.Equal(t1[i], t2[i]) {
			return false
		}
	}
	return true
}
