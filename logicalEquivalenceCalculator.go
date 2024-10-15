package main

import (
	"fmt"
	"slices"
)

func LogicalEquivalenceCalculator(expr1 string, expr2 string) (result bool, err error) {
	states1, err := CreateTruthTable(expr1)
	if err != nil {
		return false, err
	}
	fmt.Println()
	states2, err := CreateTruthTable(expr2)
	if err != nil {
		return false, err
	}
	exp1 := State("")
	for k := range states1 {
		if k == State(expr1) {
			exp1 = k
			break
		}
		if len(exp1) < len(k) {
			exp1 = k
		}
	}
	exp2 := State("")
	for k := range states2 {
		if k == State(expr2) {
			exp2 = k
			break
		}
		if len(exp2) < len(k) {
			exp2 = k
		}
	}
	if len(states1[exp1]) == len(states2[exp2]) && !slices.Equal(states1[exp1], states2[exp2]) {
		return false, nil
	}

	var minBins []Binary
	var maxBins []Binary
	if len(states1[exp1]) < len(states2[exp2]) {
		minBins = states1[exp1]
		maxBins = states2[exp2]
	} else {
		minBins = states2[exp2]
		maxBins = states1[exp1]
	}

	for i := 0; i < len(minBins); i++ {
		offset := len(maxBins) / len(minBins)
		for j := 0; j < offset; j++ {
			if minBins[i] != maxBins[(i*offset)+j] {
				return false, nil
			}
		}
	}
	return true, nil
}
