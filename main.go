package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := normalizeExpression(strings.ToLower(scanner.Text()))
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	if err != nil {
		fmt.Println("\033[91mNot a valid expression\033[97m")
		return
	}
	states := PopulatesStateBins(tokens, stateNames)
	PrintTable(stateNames, &states)
}

func normalizeExpression(expr string) string {
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
