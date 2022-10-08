package main

import (
	"fmt"
)

func main() {
	var n, m, l int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&l)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&b[i][j])
		}
	}
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
	}

	for k := 0; k < n; k++ {
		for j := 0; j < l; j++ {
			for i := 0; i < m; i++ {
				c[k][j] += a[k][i] * b[i][j]
			}
		}
	}
	for _, v := range c {
		for j := 0; j < len(v)-1; j++ {
			fmt.Printf("%d ", v[j])
		}
		fmt.Println(v[len(v)-1])
	}
}
