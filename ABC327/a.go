package main

import (
    "fmt"
    "bufio"
		"os"
		"strings"
)

func main() {
    var N int
    fmt.Scan(&N)
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), "")
    
    for i:=0;i<N-1;i++{
      if (inputs[i]=="a" && inputs[i+1]=="b" )||(inputs[i]=="b" && inputs[i+1]=="a") {
      fmt.Println("Yes")
      return
      } 
    }
    fmt.Println("No")
    
}