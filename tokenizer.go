package main

import (
	"errors"
	"fmt"
	"unicode"

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

func inputDotsForAnd(expr string) string {
	prev := expr[0]
	result := ""
	result += string(prev)
	for i := 1; i < len(expr); i++ {
		if expr[i] == ' ' {
			continue
		}
		if unicode.IsLetter(rune(expr[i])) {
			if unicode.IsLetter(rune(prev)) {
				result += "."
			}
		}
		result += string(expr[i])
		prev = expr[i]
	}
	return result
}

func FillTokensAndStates(tokens *[]Token, states *States, expr string) error {
	expr = inputDotsForAnd(expr)
	st := stack.New()
	for i := 0; i < len(expr); i++ {
		switch c := expr[i]; c {
		case '(':
			st.Push(c)
			*tokens = append(*tokens, Token{
				Value: c,
				Type:  TokenBracketOpen,
			})
		case ')':
			if st.Len() == 0 {
				return errors.New("ERROR: Invalid Syntax")
			}
			st.Pop()
			*tokens = append(*tokens, Token{
				Value: c,
				Type:  TokenBracketClose,
			})
		case '+':
			*tokens = append(*tokens, Token{
				Value: c,
				Type:  TokenOr,
			})
		case '\'':
			*tokens = append(*tokens, Token{
				Value: c,
				Type:  TokenNot,
			})
		case '.':
			*tokens = append(*tokens, Token{
				Value: c,
				Type:  TokenAnd,
			})
		default:
			if c != ' ' {
				*tokens = append(*tokens, Token{
					Value: c,
					Type:  TokenBool,
				})
				(*states)[c] = nil
			}
		}
	}
	fmt.Println(tokens)
	return nil
}
