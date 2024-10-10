package main

import (
	"strconv"
)

func GenerateTable(tokens *[]Token, states *States, stateNames []State) (table []Binary, error error) {
	numberOfStates := len(stateNames)
	numberOfRows := 2 << (numberOfStates - 1)
	table = make([]Binary, numberOfRows)
	binaries := getAllBinaryRows(numberOfRows, numberOfStates)
	for col, state := range stateNames {
		for row, binary := range binaries {
			// Need a better visualization for this one
			(*states)[state][row] = binary[col]
			if _, exists := (*states)[state+"'"]; !exists {
			}
		}
	}
	return table, nil
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
		bins := make([]Binary, 0)
		for _, c := range bin {
			switch c {
			case '0':
				bins = append(bins, 0)
			case '1':
				bins = append(bins, 1)
			}
		}
		binaryRows = append(binaryRows, bins)
	}
	return binaryRows
}
