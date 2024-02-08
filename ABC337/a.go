package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
		"strings"
)

func main() {
    var N,X,Y int
    fmt.Scan(&N)
    
    scanner := bufio.NewScanner(os.Stdin)
    
    var x,y int
    
    for i:=0; i< N;i++{
    scanner.Scan()
    inputs := strings.Split(scanner.Text(), " ")
    x, _ = strconv.Atoi(inputs[0])
    y, _ = strconv.Atoi(inputs[1])
    X=X+x
    Y=Y+y
    inputs=nil
    }

    
    if X>Y{
      fmt.Println("Takahashi")
    } else if X<Y{
      fmt.Println("Aoki")
    } else {
      fmt.Println("Draw")
    }
  
}