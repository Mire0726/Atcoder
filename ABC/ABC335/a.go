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
    inputs := strings.Split(scanner.Text(), "")
    inputs[len(inputs)-1]="4"
    
    Newinputs:=strings.Join(inputs, "")
    
    fmt.Println(Newinputs)
}