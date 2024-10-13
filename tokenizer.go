package main

import (
	"errors"
	"slices"
	"sort"
	"unicode"

	"github.com/golang-collections/collections/stack"
)

type Token struct {
	Value string
	Type  uint8
}

const (
	TokenBool uint8 = iota
	TokenNotBool
	TokenOr
	TokenAnd
	TokenBracketOpen
	TokenBracketClose
)

func GenerateTokensAndStates(expr string) (tokens []Token, stateNames []State, err error) {
	expr = normalizeExpression(expr)
	st := stack.New()
	stNames := make([]string, 0)
	if expr[0] == '+' || expr[len(expr)-1] == '+' {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}
	for i := 0; i < len(expr); i++ {
		switch c := expr[i]; c {
		case '(':
			if i != 0 &&
				(tokens[len(tokens)-1].Type == TokenBool ||
					tokens[len(tokens)-1].Type == TokenNotBool ||
					tokens[len(tokens)-1].Type == TokenOr ||
					expr[0] != '(') {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			st.Push(c)
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketOpen,
			})
		case ')':
			if st.Len() == 0 {
				return nil, nil, errors.New("ERROR: Invalid Syntax please chech the brackets properly")
			}
			if tokens[len(tokens)-1].Type == TokenOr {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			st.Pop()
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketClose,
			})
		case '+':
			if expr[i-1] == ')' || expr[i-1] == '(' {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenOr,
			})
		case '\'':
			tokens[len(tokens)-1].Value += string(c)
			tokens[len(tokens)-1].Type += TokenNotBool
		case '.':
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenAnd,
			})
		default:
			if c != ' ' {
				tokens = append(tokens, Token{
					Value: string(c),
					Type:  TokenBool,
				})
				if !slices.Contains(stNames, string(c)) {
					stNames = append(stNames, string(c))
				}
			}
		}
	}
	if st.Len() != 0 {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}
	sort.Strings(stNames)
	for _, v := range stNames {
		stateNames = append(stateNames, State(v))
	}
	numOfStates := len(stateNames)
	for i := 0; i < numOfStates; i++ {
		stateNames = append(stateNames, stateNames[i]+"'")
	}
	return tokens, stateNames, nil
}

func normalizeExpression(expr string) string {
	prev := expr[0]
	result := ""
	result += string(prev)
	for i := 1; i < len(expr); i++ {
		if expr[i] == ' ' {
			continue
		} else if unicode.IsLetter(rune(expr[i])) {
			if unicode.IsLetter(rune(prev)) || prev == '\'' {
				result += "."
			}
		}
		result += string(expr[i])
		prev = expr[i]
	}
	return result
}
