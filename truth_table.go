package main

import (
	"fmt"
	"strings"
)

func CalculateTermBinaries(terms []Term, isPos bool, states States) (States, []State) {
	numVars := len(states) / 2
	if numVars == 0 {
		return states, nil
	}
	numOfRows := 1 << numVars
	termStrings := make([]State, 0, len(terms))
	for _, term := range terms {
		termString := ""
		finalBinary := make([]Binary, numOfRows)
		first := true
		if isPos {
			for _, c := range term {
				termString += c.Value
				if c.Type != TokenAnd {
					col, ok := states[State(c.Value)]
					if !ok {
						continue
					}
					for k, b := range col {
						if first {
							finalBinary[k] = b
						} else {
							finalBinary[k] &= b
						}
					}
					first = false
				}
			}
		} else {
			for _, c := range term {
				termString += c.Value
				if c.Type != TokenOr && c.Type != TokenBracketClose && c.Type != TokenBracketOpen {
					col, ok := states[State(c.Value)]
					if !ok {
						continue
					}
					for k, b := range col {
						if first {
							finalBinary[k] = b
						} else {
							finalBinary[k] |= b
						}
					}
					first = false
				}
			}
		}
		st := State(termString)
		termStrings = append(termStrings, st)
		states[st] = finalBinary
	}
	return states, termStrings
}

func PrintTable(stateNames *[]State, termStrings *[]State, states *States) {
	if stateNames == nil || len(*stateNames) == 0 {
		return
	}
	numberOfRows := 1 << (len(*stateNames) / 2)
	for _, v := range *stateNames {
		fmt.Printf(" %v ", v)
	}
	if termStrings != nil {
		for _, v := range *termStrings {
			fmt.Printf(" %v ", v)
		}
	}
	fmt.Println()
	for row := 0; row < numberOfRows; row++ {
		for _, state := range *stateNames {
			st := (*states)[state]
			bin := st[row]
			fmt.Printf(" %v ", bin)
			printSpaces(len(state) - 1)
		}
		if termStrings != nil {
			for _, term := range *termStrings {
				st := (*states)[term]
				bin := st[row]
				fmt.Printf(" %v ", bin)
				printSpaces(len(term) - 1)
			}
		}
		fmt.Println()
	}
}

func printSpaces(number int) {
	if number > 0 {
		fmt.Print(strings.Repeat(" ", number))
	}
}

func CalculateFinalTable(termStrings *[]State, isPos bool, states States) States {
	if termStrings == nil || len(*termStrings) == 0 {
		return states
	}
	numOfRows := len(states[(*termStrings)[0]])
	finalTable := make([]Binary, numOfRows)
	finalString := State("")
	for i, term := range *termStrings {
		col := states[term]
		if isPos {
			for j, b := range col {
				if i == 0 {
					finalTable[j] = b
				} else {
					finalTable[j] |= b
				}
			}
			if i != 0 {
				finalString += "+" + term
			} else {
				finalString += term
			}
		} else {
			for j, b := range col {
				if i == 0 {
					finalTable[j] = b
				} else {
					finalTable[j] &= b
				}
			}
			finalString += term
		}
	}
	states[finalString] = finalTable
	*termStrings = append(*termStrings, finalString)
	return states
}

func evaluate(tokens []Token, stateNames []State) (States, []State, []Binary) {
	terms, isPos := ParseTerms(&tokens)
	states := PopulatesStateBins(tokens, stateNames)
	states, termStrings := CalculateTermBinaries(terms, isPos, states)
	states = CalculateFinalTable(&termStrings, isPos, states)
	finalKey := termStrings[len(termStrings)-1]
	finalBinary := states[finalKey]
	return states, termStrings, finalBinary
}

func CreateTruthTable(expr string, flag uint8) (States, error) {
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	if err != nil {
		return nil, err
	}
	states, termStrings, _ := evaluate(tokens, stateNames)
	if flag == Print {
		PrintTable(&stateNames, &termStrings, &states)
	}
	return states, nil
}
