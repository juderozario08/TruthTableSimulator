package main

import (
	"errors"
	"slices"

	"github.com/golang-collections/collections/stack"
)

type Token struct {
	Value byte
	Type  int
}

const (
	TokenBool = iota
	TokenOr
	TokenAnd
	TokenBracketOpen
	TokenBracketClose
	TokenNot
)

func GenerateTokensAndStates(expr string) (tokens []Token, stateNames []State, err error) {
	st := stack.New()
	stateNames = make([]State, 0)
	for i := 0; i < len(expr); i++ {
		switch c := expr[i]; c {
		case '(':
			st.Push(c)
			tokens = append(tokens, Token{
				Value: c,
				Type:  TokenBracketOpen,
			})
		case ')':
			if st.Len() == 0 {
				return nil, nil, errors.New("ERROR: Invalid Syntax")
			}
			st.Pop()
			tokens = append(tokens, Token{
				Value: c,
				Type:  TokenBracketClose,
			})
		case '+':
			tokens = append(tokens, Token{
				Value: c,
				Type:  TokenOr,
			})
		case '\'':
			tokens = append(tokens, Token{
				Value: c,
				Type:  TokenNot,
			})
		case '.':
			tokens = append(tokens, Token{
				Value: c,
				Type:  TokenAnd,
			})
		default:
			if c != ' ' {
				tokens = append(tokens, Token{
					Value: c,
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
