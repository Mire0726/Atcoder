package main

import (
  "fmt"
  "unicode"
  )

func main() {
    var N string
    fmt.Scan(&N)
    sli:=[]rune(N)
    
    // 1文字だけの場合のチェック
    if len(sli) == 1 {
        if unicode.IsUpper(sli[0]) {
            fmt.Println("Yes")
        } else {
            fmt.Println("No")
        }
        return
    }

    // 最初の文字が大文字かつ残りが小文字かチェック
    if unicode.IsUpper(sli[0]) {
        allLower := true
        for i := 1; i < len(sli); i++ {
            if !unicode.IsLower(sli[i]) {
                allLower = false
                break
            }
        }
        if allLower {
            fmt.Println("Yes")
            return
        }
    }
    fmt.Println("No")
    
}
    


