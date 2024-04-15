package main
import "fmt"

func floordiv(x int, y int) int {
	r := (x % y + y) % y
	return (x - r) / y
}

func main() {
	var A, M, L, R int; fmt.Scanf("%d %d %d %d", &A, &M, &L, &R);
	L, R = L-A, R-A

	fmt.Println(floordiv(R, M) - floordiv(L-1, M))
}