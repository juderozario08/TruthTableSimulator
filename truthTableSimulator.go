package main

import (
	"fmt"
	"strconv"
)

func PopulatesStateBins(tokens []Token, stateNames []State) (states States) {
	numberOfStates := len(stateNames)
	numberOfRows := 2 << (numberOfStates/2 - 1)
	fmt.Println(stateNames)
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
		if len(bin) < numberOfStates/2 {
			for j := 0; j < (numberOfStates/2)-len(bin); j++ {
				remainingBits += "0"
			}
		}
		bin = remainingBits + bin
		bins := make([]Binary, numberOfStates)
		for k, c := range bin {
			switch c {
			case '0':
				bins[k] = 0
				bins[k+(numberOfStates/2)] = 1
			case '1':
				bins[k] = 1
				bins[k+(numberOfStates/2)] = 0
			}
		}
		binaryRows = append(binaryRows, bins)
	}
	return binaryRows
}

func PrintTable(stateNames []State, states *States) {
	numberOfRows := 2 << (len(stateNames)>>1 - 1)
	fmt.Println(numberOfRows)
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
