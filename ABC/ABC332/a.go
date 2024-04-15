package main

import (
    "fmt"
)

func main() {
    var N,S,Q int
    fmt.Scan(&N,&S,&Q)
  
    x:=0
    
    for i:=0;i<N;i++{
      var a,b int
      fmt.Scan(&a,&b)
      
      x=x+a * b
    }
    
    if x<S{
        fmt.Println(x+Q)
      } else{
        fmt.Println(x)
      }
    
}