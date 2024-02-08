package main

import (
    "fmt"
    "strings"
    "bufio"
		"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), "")

    
    ans :=strings.Join(inputs," ")
    fmt.Println(ans)
    
    
}