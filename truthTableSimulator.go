package main

import (
	"strconv"
)

func GenerateTable(tokens *[]Token, states *States) (table []int, error error) {
	numberOfRows := 2 << (len(*states) - 1)
	table = make([]int, numberOfRows)
	strBin := normalizeBinaryLength(numberOfRows, len(*states))
	return table, nil
}

func normalizeBinaryLength(numberOfRows int, numberOfStates int) []string {
	strBin := make([]string, 0)
	for i := 0; i < numberOfRows; i++ {
		bin := strconv.FormatInt(int64(i), 2)
		remainingBits := ""
		if len(bin) < numberOfStates {
			for j := 0; j < numberOfStates-len(bin); j++ {
				remainingBits += "0"
			}
		}
		bin = remainingBits + bin
		strBin = append(strBin, bin)
	}
	return strBin
}
