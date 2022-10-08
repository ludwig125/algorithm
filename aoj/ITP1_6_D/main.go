package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([]int, m)
	for j := 0; j < m; j++ {
		fmt.Scan(&b[j])
	}

	for i := 0; i < n; i++ {
		ans := 0
		for j := 0; j < m; j++ {
			ans += a[i][j] * b[j]
		}
		fmt.Println(ans)
	}
}
