package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NoPrint uint8 = iota
	Print
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please select one from the following:\n1. Truth Table Simulator\n2. Logical Equivalence Calculator\n")
	if !scanner.Scan() {
		return
	}
	choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Printf("\033[91m%v\033[97m\n", err.Error())
		return
	}
	if choice > 2 || choice < 1 {
		fmt.Println("\033[91mPlease type either 1 or 2\033[97m")
		return
	}
	if choice == 1 {
		fmt.Println("Enter an SOP (minterms) or POS (maxterms) function: ")
		if !scanner.Scan() {
			return
		}
		expr := strings.ToLower(strings.TrimSpace(scanner.Text()))
		_, err := CreateTruthTable(expr, Print)
		if err != nil {
			fmt.Printf("\033[91m%v\033[97m\n", err.Error())
			return
		}
	} else {
		fmt.Println("Enter the 1st SOP (minterms) or POS (maxterms) function: ")
		if !scanner.Scan() {
			return
		}
		expr1 := strings.ToLower(strings.TrimSpace(scanner.Text()))
		fmt.Println("Enter the 2nd SOP (minterms) or POS (maxterms) function: ")
		if !scanner.Scan() {
			return
		}
		expr2 := strings.ToLower(strings.TrimSpace(scanner.Text()))
		result, err := LogicalEquivalenceCalculator(expr1, expr2, Print)
		if err != nil {
			fmt.Printf("\033[91m%v\033[97m\n", err.Error())
			return
		}
		if result {
			fmt.Println("\033[92mThese 2 logical statements are equivalent\033[97m")
		} else {
			fmt.Println("\033[91mThe 2 logical statements are not equivalent\033[97m")
		}
	}
}
