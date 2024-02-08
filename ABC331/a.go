package main

import (
    "fmt"
)

func main() {
    var M,D,y,m,d int
    fmt.Scan(&M,&D,&y,&m,&d)

    
    if m==M &&d==D{
      fmt.Println(y+1,1,1)
    } else if d==D{
      fmt.Println(y,m+1,1)
    } else{
      fmt.Println(y,m,d+1)
    }
   
}