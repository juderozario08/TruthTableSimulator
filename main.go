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
	fmt.Println("Enter an SOP(maxterms) or POS(minterms) function: ")
	scanner.Scan()
	input := strings.ToLower(scanner.Text())
	expr := normalizeExpression(input)
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	if err != nil {
		fmt.Println("\033[91mNot a valid expression\033[97m")
		return
	}
	terms, isPos := ParseTerms(&tokens)
	fmt.Println(terms)
	states := PopulatesStateBins(tokens, stateNames)
	states = CalculateTermBinaries(terms, isPos, states)
	PrintTable(stateNames, &states)
}

func normalizeExpression(expr string) string {
	prev := expr[0]
	result := ""
	result += string(prev)
	for i := 1; i < len(expr); i++ {
		if expr[i] == ' ' {
			continue
		} else if unicode.IsLetter(rune(expr[i])) {
			if unicode.IsLetter(rune(prev)) {
				result += "."
			}
		}
		result += string(expr[i])
		prev = expr[i]
	}
	return result
}
