package main

import (
	"fmt"
)

func CalculateTermBinaries(terms []Term, isPos bool, states States) (sts States, termString []State) {
	numOfRows := 2 << (len(states)/2 - 1)
	termStrings := make([]State, 0)
	for _, term := range terms {
		termString := ""
		finalBinary := make([]Binary, numOfRows)
		if isPos {
			for j, c := range term {
				termString += c.Value
				if c.Type != TokenAnd {
					for k, b := range states[State(c.Value)] {
						if j == 0 {
							finalBinary[k] |= b
						} else {
							finalBinary[k] &= b
						}
					}
				}
			}
		} else {
			for j, c := range term {
				termString += c.Value
				if c.Type != TokenOr && c.Type != TokenBracketClose && c.Type != TokenBracketOpen {
					for k, b := range states[State(c.Value)] {
						if j == 0 {
							finalBinary[k] |= b
						} else {
							finalBinary[k] |= b
						}
					}
				}
			}
		}
		termStrings = append(termStrings, State(termString))
		states[State(termString)] = finalBinary
	}
	return states, termStrings
}

func PrintTable(stateNames *[]State, termStrings *[]State, states *States) {
	numberOfRows := 2 << (len(*stateNames)>>1 - 1)
	for _, v := range *stateNames {
		fmt.Printf(" %v ", v)
	}
	for _, v := range *termStrings {
		fmt.Printf(" %v ", v)
	}
	fmt.Println()
	for row := 0; row < numberOfRows; row++ {
		for _, state := range *stateNames {
			st := (*states)[state]
			bin := st[row]
			fmt.Printf(" %v ", bin)
			printSpaces(len(state) - 1)
		}
		for _, term := range *termStrings {
			st := (*states)[term]
			bin := st[row]
			fmt.Printf(" %v ", bin)
			printSpaces(len(term) - 1)
		}
		fmt.Println()
	}
}

func printSpaces(number int) {
	for i := 0; i < number; i++ {
		fmt.Print(" ")
	}
}

func CalculateFinalTable(termStrings *[]State, isPos bool, states States) States {
	finalTable := make([]Binary, 0)
	finalString := State("")
	for i, term := range *termStrings {
		if isPos {
			for j, b := range states[term] {
				if i == 0 {
					finalTable = append(finalTable, b)
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
			for j, b := range states[term] {
				if i == 0 {
					finalTable = append(finalTable, b)
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

func CreateTruthTable(expr string, flag uint8) (States, error) {
	tokens, stateNames, err := GenerateTokensAndStates(expr)
	if err != nil {
		return nil, err
	}
	terms, isPos := ParseTerms(&tokens)
	states := PopulatesStateBins(tokens, stateNames)
	states, termStrings := CalculateTermBinaries(terms, isPos, states)
	states = CalculateFinalTable(&termStrings, isPos, states)
	if flag == Print {
		PrintTable(&stateNames, &termStrings, &states)
	}
	return states, nil
}
