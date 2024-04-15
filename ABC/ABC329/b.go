package main

import (
    "fmt"
    "sort"
)

func main() {
    var N int
    fmt.Scan(&N)

    arr := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&arr[i])
    }
    sort.Ints(arr) // 配列を昇順にソート

    // 最大値を見つける
    maxVal := arr[N-1]

    // 最大値でない中で最大の値を見つける
    for i := N - 2; i >= 0; i-- {
        if arr[i] != maxVal {
            fmt.Println(arr[i]) // 2番目に大きい値を出力
            return
        }
    }
}