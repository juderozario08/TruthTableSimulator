package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an SOP(maxterms) or POS(minterms) function: ")
	scanner.Scan()
	expr := strings.ToLower(scanner.Text())
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	if err != nil {
		fmt.Printf("\033[91m%v\033[97m\n", err.Error())
		return
	}
	terms, isPos := ParseTerms(&tokens)
	states := PopulatesStateBins(tokens, stateNames)
	states = CalculateTermBinaries(terms, isPos, states)
	stateStrings := make([]string, 0)
	for k := range states {
		stateStrings = append(stateStrings, string(k))
	}
	// TODO: Build custom string sorting algorithm
	sort.Strings(stateStrings)
	for _, st := range stateStrings {
		fmt.Println(st, states[State(st)])
	}
	// Come up with a better way to print the table
	// PrintTable(&stateNames, &states)
}
