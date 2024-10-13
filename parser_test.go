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

func TestParser(t *testing.T) {
	tests := []ParseTermsTest{
		{
			ExpectedTerms: []Term{
				[]Token{
					{Value: "a", Type: TokenBool},
					{Value: ".", Type: TokenAnd},
					{Value: "b", Type: TokenBool},
					{Value: ".", Type: TokenAnd},
					{Value: "c", Type: TokenBool},
				},
				[]Token{
					{Value: "c", Type: TokenBool},
					{Value: ".", Type: TokenAnd},
					{Value: "d", Type: TokenBool},
					{Value: ".", Type: TokenAnd},
					{Value: "e", Type: TokenBool},
				},
			},
			ExpectedIsPOS: true,
			Question:      "abc + cde",
		},
		{
			ExpectedTerms: []Term{
				[]Token{
					{Value: "(", Type: TokenBracketOpen},
					{Value: "a", Type: TokenBool},
					{Value: "+", Type: TokenOr},
					{Value: "b", Type: TokenBool},
					{Value: "+", Type: TokenOr},
					{Value: "c", Type: TokenBool},
					{Value: ")", Type: TokenBracketClose},
				},
				[]Token{
					{Value: "(", Type: TokenBracketOpen},
					{Value: "c", Type: TokenBool},
					{Value: "+", Type: TokenOr},
					{Value: "d", Type: TokenBool},
					{Value: "+", Type: TokenOr},
					{Value: "e", Type: TokenBool},
					{Value: ")", Type: TokenBracketClose},
				},
			},
			ExpectedIsPOS: false,
			Question:      "(a+b+c)(c+d+e)",
		},
	}
	for _, test := range tests {
		tokens, _, _ := GenerateTokensAndStates(test.Question)
		terms, isPOS := ParseTerms(&tokens)
		if !termEqual(terms, test.ExpectedTerms) {
			t.Errorf("Expected: %v\nGot: %v", test.ExpectedTerms, terms)
		}
		if isPOS != test.ExpectedIsPOS {
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
