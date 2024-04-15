package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scan(&N)

    arr := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&arr[i])
    }

    result := true // 最初に真を仮定
    for n := 1; n < N; n++ { // 最初の要素と比較するため、インデックスは1から始める
        if arr[0] != arr[n] { // 最初の要素と異なる要素があれば、全て等しくない
            result = false
            break
        }
    }

    if result {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
