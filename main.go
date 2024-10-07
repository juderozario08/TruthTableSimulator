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
	tokens, err := GenerateTokens(expr)
	if err != nil {
		fmt.Println("\033[91mNot a valid expression\033[97m")
		return
	}
	fmt.Println(tokens)
}
