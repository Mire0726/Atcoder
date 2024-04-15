package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	fmt.Scan(&n)

	s := strconv.FormatInt(n, 2)
	r := []rune(s)
	t := []rune("0")[0]
	count := 0

	for i := len(r) - 1; i >= 0; i-- {
		if r[i] != t {
			break
		}
		count++
	}

	fmt.Println(count)
}