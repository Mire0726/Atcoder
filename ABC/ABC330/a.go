package main

import (
    "fmt"
)

func main() {
    var N,L int
    fmt.Scan(&N,&L)
 
    arr:=make([]int,N)
    cnt:=0
    
    for i:=0;i<N;i++{
      fmt.Scan(&arr[i])
      
      if arr[i]>=L{
        cnt++
      }
    }
    fmt.Println(cnt)
    
    
}