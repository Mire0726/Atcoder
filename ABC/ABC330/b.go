package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var N,L,R int
    fmt.Scan(&N,&L,&R)
 
    arr:=make([]int,N)
    ans:=make([]int,0)
    
    for i:=0;i<N;i++{
      fmt.Scan(&arr[i])
      
      if arr[i]<L{
       ans=append(ans,L)
     } else if arr[i]>=R{
       ans =append(ans,R)
     } else {
       ans=append(ans,arr[i])
     }
}
var Ans[] string
// スライス内の各数字を文字列に変換して結合
    for _, num := range ans {
        Ans = append(Ans, strconv.Itoa(num))
    }
    final:=strings.Join(Ans," ")

  fmt.Println(final)
}