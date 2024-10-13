package main

import (
	"slices"
	"testing"
)

type TokenizerTest struct {
	ExpectedToken      []Token
	ExpectedStateNames []State
	Question           string
}

func TestTokenizerAndStateNames(t *testing.T) {
	tests := []TokenizerTest{
		{
			ExpectedToken: []Token{
				{Type: TokenBracketOpen, Value: "("},
				{Type: TokenBool, Value: "a"},
				{Type: TokenOr, Value: "+"},
				{Type: TokenBool, Value: "c"},
				{Type: TokenBracketClose, Value: ")"},
				{Type: TokenBracketOpen, Value: "("},
				{Type: TokenBool, Value: "d"},
				{Type: TokenOr, Value: "+"},
				{Type: TokenNotBool, Value: "f'"},
				{Type: TokenBracketClose, Value: ")"},
			},
			ExpectedStateNames: []State{"a", "c", "d", "f", "a'", "c'", "d'", "f'"},
			Question:           "(a + c)(d + f')", // SOP test
		},
		{
			ExpectedToken: []Token{
				{Type: TokenBracketOpen, Value: "("},
				{Type: TokenBool, Value: "a"},
				{Type: TokenOr, Value: "+"},
				{Type: TokenBool, Value: "c"},
				{Type: TokenBracketClose, Value: ")"},
				{Type: TokenBracketOpen, Value: "("},
				{Type: TokenBool, Value: "b"},
				{Type: TokenOr, Value: "+"},
				{Type: TokenNotBool, Value: "f'"},
				{Type: TokenBracketClose, Value: ")"},
			},
			ExpectedStateNames: []State{"a", "b", "c", "f", "a'", "b'", "c'", "f'"},
			Question:           "(a + c)(b + f')", // SOP test
		},
		{
			ExpectedToken: []Token{
				{Type: TokenBool, Value: "a"},
				{Type: TokenAnd, Value: "."},
				{Type: TokenBool, Value: "b"},
				{Type: TokenAnd, Value: "."},
				{Type: TokenBool, Value: "c"},
				{Type: TokenOr, Value: "+"},
				{Type: TokenNotBool, Value: "c'"},
				{Type: TokenAnd, Value: "."},
				{Type: TokenBool, Value: "e"},
				{Type: TokenAnd, Value: "."},
				{Type: TokenBool, Value: "f"},
			},
			ExpectedStateNames: []State{"a", "b", "c", "e", "f", "a'", "b'", "c'", "e'", "f'"},
			Question:           "abc + c'ef", // POS test
		},
		{
			ExpectedToken:      nil,
			ExpectedStateNames: nil,
			Question:           "((abc)", // nil test
		},
		{
			ExpectedToken:      nil,
			ExpectedStateNames: nil,
			Question:           "(a + b + c)+a.b.c", // SOP POS err test
		},
		{
			ExpectedToken:      nil,
			ExpectedStateNames: nil,
			Question:           "abc(c + b + e)", // POS SOP err test
		},
	}
	for _, test := range tests {
		tokens, stateNames, _ := GenerateTokensAndStates(test.Question)
		if !slices.Equal(tokens, test.ExpectedToken) {
			t.Errorf("Expected: %v,\nGot: %v\n", test.ExpectedToken, tokens)
		}
		if !slices.Equal(stateNames, test.ExpectedStateNames) {
			t.Errorf("Expected: %v,\nGot: %v\n", test.ExpectedStateNames, stateNames)
		}
	}
}
