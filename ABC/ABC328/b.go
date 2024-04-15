package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scan(&N)

    D := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&D[i])
    }

    count := 0
    for month := 1; month <= N; month++ {
        // 各月について、日数がその月の数字を繰り返す形かどうかをチェック
        for day := 1; day <= D[month-1]; day++ {
            if isZorome(month, day) {
                count++
            }
        }
    }

    fmt.Println(count)
}

// isZorome 関数は、与えられた月と日がゾロ目の条件を満たすかどうかをチェックします。
func isZorome(month, day int) bool {
    // 月が一桁で月と日が同じ場合、ゾロ目とみなす
    if month < 10 && month == day {
        return true
    }
    // 月が1〜9で、日がその月の数字を繰り返す場合（1月11日など）
    if month < 10 && (day == month*10+month) {
        return true
    }
    // 11月1日、2月22日などのケースをチェック
    if day < 10 && month == day*10+day{
        return true
    }
    // 11月11日、22月22日などのケースをチェック
    if day ==month && day%11==0 {
        return true
    }
    return false
}
