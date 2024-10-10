package main

import (
	"fmt"
	"strconv"
)

func GenerateTable(tokens []Token, stateNames []State) (states States) {
	numberOfStates := len(stateNames)
	numberOfRows := 2 << (numberOfStates/2 - 1)
	states = make(States)
	binaries := getAllBinaryRows(numberOfRows, numberOfStates)
	for row, bins := range binaries {
		for col, bin := range bins {
			_, exists := states[stateNames[col]]
			if !exists {
				states[stateNames[col]] = make([]Binary, numberOfRows)
			}
			states[stateNames[col]][row] = bin
		}
	}
	return states
}

func getAllBinaryRows(numberOfRows int, numberOfStates int) [][]Binary {
	binaryRows := make([][]Binary, 0)
	for i := 0; i < numberOfRows; i++ {
		bin := strconv.FormatInt(int64(i), 2) // Get binary
		remainingBits := ""
		if len(bin) < numberOfStates {
			for j := 0; j < numberOfStates-len(bin); j++ {
				remainingBits += "0"
			}
		}
		bin = remainingBits + bin
		bins := make([]Binary, numberOfStates)
		for k, c := range bin {
			switch c {
			case '0':
				bins[k] = 0
			case '1':
				bins[k] = 1
			}
		}
		binaryRows = append(binaryRows, bins)
	}
	return binaryRows
}

func PrintTable(numberOfRows int, stateNames []State, states States) {
	numberOfStates := len(stateNames)
	fmt.Print("-")
	for i := 0; i < numberOfStates; i++ {
		fmt.Print("---")
	}
	fmt.Print("-\n")
	fmt.Print("|")
	for _, state := range stateNames {
		fmt.Print(" " + state + " ")
	}
	fmt.Print("|\n")
	fmt.Print("-")
	for i := 0; i < numberOfStates; i++ {
		fmt.Print("---")
	}
	fmt.Print("-\n")
	for i := 0; i < numberOfRows; i++ {
		fmt.Print("|")
		for _, v := range states {
			fmt.Printf(" %v ", v[i])
		}
		fmt.Print("|\n")
	}
	fmt.Print("-")
	for i := 0; i < numberOfStates; i++ {
		fmt.Print("---")
	}
	fmt.Print("-\n")
}
