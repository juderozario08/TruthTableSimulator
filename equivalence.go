package main

import (
	"fmt"
	"slices"
	"sort"
)

func LogicalEquivalenceCalculator(expr1 string, expr2 string, flag uint8) (result bool, err error) {
	tokens1, stateNames1, err := GenerateTokensAndStates(expr1)
	if err != nil {
		return false, err
	}
	tokens2, stateNames2, err := GenerateTokensAndStates(expr2)
	if err != nil {
		return false, err
	}

	if flag == Print {
		states1, termStrings1, _ := evaluate(tokens1, stateNames1)
		PrintTable(&stateNames1, &termStrings1, &states1)

		fmt.Println()

		states2, termStrings2, _ := evaluate(tokens2, stateNames2)
		PrintTable(&stateNames2, &termStrings2, &states2)
	}

	varMap := make(map[string]bool)
	for _, st := range stateNames1[:len(stateNames1)/2] {
		varMap[string(st)] = true
	}
	for _, st := range stateNames2[:len(stateNames2)/2] {
		varMap[string(st)] = true
	}

	unionVars := make([]string, 0, len(varMap))
	for v := range varMap {
		unionVars = append(unionVars, v)
	}
	sort.Strings(unionVars)

	allStateNames := make([]State, 0, len(unionVars)*2)
	for _, v := range unionVars {
		allStateNames = append(allStateNames, State(v))
	}
	for _, v := range unionVars {
		allStateNames = append(allStateNames, State(v+"'"))
	}

	_, _, final1 := evaluate(tokens1, allStateNames)
	_, _, final2 := evaluate(tokens2, allStateNames)

	return slices.Equal(final1, final2), nil
}
