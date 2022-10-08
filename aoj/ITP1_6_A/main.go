package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	for i := n - 1; i > 0; i-- {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println(a[0])
}
