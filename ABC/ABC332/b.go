package main

import (
    "fmt"
)

func main() {
  var K,G,M int
  fmt.Scan(&K,&G,&M)
  g:=0
  m:=0

  
  for i:=0;i<K;i++{
    if g==G{
      g=0
    } else if m==0{
      m=m+M
    } else{
      if G-g>m{
        g=g+m
        m=0
      } else{
        a:=G-g
        g=G
        m=m-a
      }
    }


  }
  
  fmt.Println(g,m)
}