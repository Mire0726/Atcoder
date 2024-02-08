package main

import (
    "fmt"
    "golang.org/x/exp/slices"
)

func main() {
    var S,T string
    fmt.Scan(&S,&T)

  arr :=[]string{"AC","AD","BD","BE","CE","CA","DA","DB","EB","EC"}
  
   if slices.Contains(arr, S) && slices.Contains(arr, T){
     fmt.Println("Yes")
   } else if slices.Contains(arr, S)==false && slices.Contains(arr, T)==false{
     fmt.Println("Yes")
   } else{
     fmt.Println("No")
     
   }
    
    
}