package main

import (
    "fmt"
		"bufio"
		"os"
		"strings"
		)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), ".")
    
    fmt.Println(inputs[len(inputs)-1])
}