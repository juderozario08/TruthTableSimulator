package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := strings.ToLower(scanner.Text())
	tokens := make([]Token, 0)
	states := make(States)
	err := FillTokensAndStates(&tokens, &states, expr)
	if err != nil {
		fmt.Println("\033[91mNot a valid expression\033[97m")
		return
	}
	GenerateTable(&tokens, &states)
}
