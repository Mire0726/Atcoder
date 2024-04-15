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
    inputs := strings.Split(scanner.Text(), "")
    results := 0
    
    if len(inputs) == 1 {
        results = 0
    } else {
        for i := 0; i < len(inputs)-1; i++ {
            if inputs[i] == "C" && inputs[i+1] == "B" {
                results++
            } else if inputs[i] == "B" && inputs[i+1] == "A" {
                results++
            } else if inputs[i] == "C" && inputs[i+1] == "A" {
                results++
            }
        }
    }
    
    if results > 0 {
        fmt.Println("No")
    } else {
        fmt.Println("Yes")
    }
}