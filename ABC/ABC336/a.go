
package main

import (
    "fmt"
    "strings"
)

//スペースか行区切りで複数ぎょう配列に入れる
func main() {
   var x int
   fmt.Scan(&x)
   arr:=[]string{"L"}
   
   for i:=0;i<x;i++{
     arr = append(arr,"o")
   }
   arr=append(arr,"ng")
   
   fmt.Println(strings.Join(arr,""))
}