package main

import (
    "fmt"
)


func main() {
  var X,Y int
  fmt.Scan(&X,&Y)
  
  if (X-Y<=3 && X-Y>=0)||(X-Y<0 && X-Y>=-2){
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
     
}