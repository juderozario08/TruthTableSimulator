package main

import (
	"fmt"
)

func CalculateTermBinaries(terms []Term, isPos bool, states States) States {
	numOfRows := 2 << (len(states)/2 - 1)
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
		states[State(termString)] = finalBinary
	}
	return states
}

func PrintTable(stateNames *[]State, states *States) {
	numberOfRows := 2 << (len(*stateNames)>>1 - 1)
	for _, v := range *stateNames {
		if len(v) == 2 {
			fmt.Printf("%v  ", v)
		} else {
			fmt.Printf("%v   ", v)
		}
	}
	fmt.Println()
	for row := 0; row < numberOfRows; row++ {
		for _, state := range *stateNames {
			fmt.Printf("%v   ", (*states)[state][row])
		}
		fmt.Println()
	}
}

func CalculateFinalTable(terms []Term, isPos bool, states States) States {
	return states
}
