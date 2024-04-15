package main

import ("fmt"
        "strings"
        "strconv")

func main(){
  var N int
  fmt.Scan(&N)
  
  arr:=make([]string,N)
  for i:=0;i<N;i++{
    n:=strconv.Itoa(N)
    arr=append(arr,n)
  }
  var result string
  result = strings.Join(arr,"")
  fmt.Println(result)
}