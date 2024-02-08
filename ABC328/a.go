package main

import (
    "fmt"
)

func main() {
    var N,X int
    fmt.Scan(&N,&X)

    arr:=make([]int,N)
    cnt:=0
    
    for i:=0;i<N;i++{
      fmt.Scan(&arr[i])
      if arr[i]<=X{
        cnt =cnt+arr[i]
      }
    }
    
    fmt.Println(cnt)
    
    
}