package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an SOP(maxterms) or POS(minterms) function: ")
	scanner.Scan()
	expr := strings.ToLower(scanner.Text())
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	fmt.Println(tokens)
	if err != nil {
		fmt.Printf("\033[91m%v\033[97m\n", err.Error())
		return
	}
	terms, isPos := ParseTerms(&tokens)
	fmt.Println(terms)
	fmt.Println(terms)
	states := PopulatesStateBins(tokens, stateNames)
	states = CalculateTermBinaries(terms, isPos, states)
	// PrintTable(stateNames, &states)
}
