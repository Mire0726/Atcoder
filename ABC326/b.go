package main

import (
	"fmt"
)

func main() {
  var N int
  fmt.Scan(&N)
  
  for i:=0;i<=919;i++{
    num:=N+i
    if isLike(num){
      fmt.Println(num)
      break
    }
  }
}

func isLike(n int)bool{
  a:=n%100
  b:=(n-a)/100
  c:=a%10
  d:=(a-c)/10
  if b *d ==c{
    return true
  }else {
    return false
  }
}