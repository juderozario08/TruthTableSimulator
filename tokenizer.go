package main

import (
	"errors"
	"slices"

	"github.com/golang-collections/collections/stack"
)

type Token struct {
	Value string
	Type  int
}

const (
	TokenBool = iota
	TokenNotBool
	TokenOr
	TokenAnd
	TokenBracketOpen
	TokenBracketClose
)

func GenerateTokensAndStates(expr string) (tokens []Token, stateNames []State, err error) {
	// TODO: Use a counter instead of a stack later on cause it is more memory efficient
	st := stack.New()
	stateNames = make([]State, 0)
	for i := 0; i < len(expr); i++ {
		switch c := expr[i]; c {
		case '(':
			st.Push(c)
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketOpen,
			})
		case ')':
			if st.Len() == 0 {
				return nil, nil, errors.New("ERROR: Invalid Syntax")
			}
			st.Pop()
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketClose,
			})
		case '+':
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
				if !slices.Contains(stateNames, State(c)) {
					stateNames = append(stateNames, State(c))
				}
			}
		}
	}
	if st.Len() != 0 {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}
	numOfStates := len(stateNames)
	for i := 0; i < numOfStates; i++ {
		stateNames = append(stateNames, stateNames[i]+"'")
	}
	return tokens, stateNames, nil
}
