package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	terms := Expression{
		Term: make([]any, 0),
		Type: NEITHER,
	}
	if !ValidateExpressionAndParseTerms(expr, &terms) {
		fmt.Println("\033[91mYou did not give me a proper Boolean Expression")
		fmt.Print("\033[97m")
	}
}
