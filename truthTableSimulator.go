package main

import (
	"fmt"
)

func CalculateTermBinaries(terms []Term, isPos bool, states States) States {
	numOfRows := 2 << (len(states)/2 - 1)
	for i, term := range terms {
		if isPos {
			for _, c := range term {
				finalBinary := make([]Binary, numOfRows)
				if c.Type != TokenAnd {
				}
			}
		} else {
		}
	}
	return states
}

func PrintTable(stateNames []State, states *States) {
	numberOfRows := 2 << (len(stateNames)>>1 - 1)
	for _, v := range stateNames {
		if len(v) == 2 {
			fmt.Printf("%v  ", v)
		} else {
			fmt.Printf("%v   ", v)
		}
	}
	fmt.Println()
	for row := 0; row < numberOfRows; row++ {
		for _, state := range stateNames {
			fmt.Printf("%v   ", (*states)[state][row])
		}
		fmt.Println()
	}
}
