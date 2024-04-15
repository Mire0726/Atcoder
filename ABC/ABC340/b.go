package main

import (
    "fmt"
    "bufio"
		"os"
		"strings"
		"strconv"
)

func main() {
    var Q int
    fmt.Scan(&Q)
    q :=make([]string,0)
    scanner := bufio.NewScanner(os.Stdin)
    
    for i:=0;i<Q;i++{
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), " ")
    
    if inputs[0]=="1"{
      q=append(q,inputs[1])
    } else {
      x,_:=strconv.Atoi(inputs[1])
      fmt.Println(q[len(q)-x])
    }
    }
    
}