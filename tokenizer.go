package main

import (
	"errors"
	"slices"
	"sort"
	"strings"
	"unicode"
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
	if len(expr) == 0 {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}
	if expr[0] == '+' || expr[len(expr)-1] == '+' || expr[0] == '.' || expr[len(expr)-1] == '.' {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}

	bracketDepth := 0
	stNames := make([]string, 0)

	for i := 0; i < len(expr); i++ {
		switch c := expr[i]; c {
		case '(':
			if i != 0 &&
				(tokens[len(tokens)-1].Type == TokenBool ||
					tokens[len(tokens)-1].Type == TokenNotBool ||
					tokens[len(tokens)-1].Type == TokenOr ||
					tokens[len(tokens)-1].Type == TokenAnd ||
					expr[0] != '(') {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			bracketDepth++
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketOpen,
			})
		case ')':
			if bracketDepth == 0 {
				return nil, nil, errors.New("ERROR: Invalid Syntax please check the brackets properly")
			}
			if len(tokens) > 0 && (tokens[len(tokens)-1].Type == TokenOr ||
				tokens[len(tokens)-1].Type == TokenBracketOpen ||
				tokens[len(tokens)-1].Type == TokenAnd) {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			bracketDepth--
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBracketClose,
			})
		case '+':
			if i == 0 || i == len(expr)-1 || expr[i-1] == ')' || expr[i-1] == '(' || expr[i-1] == '+' || expr[i-1] == '.' {
				return nil, nil, errors.New("ERROR: Function does not follow SOP or POS format")
			}
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenOr,
			})
		case '\'':
			if len(tokens) == 0 || tokens[len(tokens)-1].Type != TokenBool {
				return nil, nil, errors.New("ERROR: Invalid Syntax")
			}
			tokens[len(tokens)-1].Value += string(c)
			tokens[len(tokens)-1].Type = TokenNotBool
		case '.':
			if i == 0 || i == len(expr)-1 || expr[i-1] == '(' || expr[i-1] == '+' || expr[i-1] == '.' {
				return nil, nil, errors.New("ERROR: Invalid Syntax")
			}
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenAnd,
			})
		default:
			if !unicode.IsLetter(rune(c)) {
				return nil, nil, errors.New("ERROR: Invalid Syntax")
			}
			tokens = append(tokens, Token{
				Value: string(c),
				Type:  TokenBool,
			})
			if !slices.Contains(stNames, string(c)) {
				stNames = append(stNames, string(c))
			}
		}
	}

	if bracketDepth != 0 || len(stNames) == 0 {
		return nil, nil, errors.New("ERROR: Invalid Syntax")
	}

	sort.Strings(stNames)
	stateNames = make([]State, 0, len(stNames)*2)
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
	expr = strings.ReplaceAll(expr, " ", "")
	if len(expr) == 0 {
		return ""
	}
	var sb strings.Builder
	prev := rune(expr[0])
	sb.WriteRune(prev)
	for _, c := range expr[1:] {
		if unicode.IsLetter(c) && (unicode.IsLetter(prev) || prev == '\'') {
			sb.WriteRune('.')
		}
		sb.WriteRune(c)
		prev = c
	}
	return sb.String()
}
