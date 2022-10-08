// https://qiita.com/drken/items/fd4e5e3630d0f5859067
package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}

	//fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//fmt.Println(a)

	A := 0
	B := 0
	for i := 0; i < len(a); i += 2 {
		A += a[i]
		if i+1 < len(a) {
			B += a[i+1]
		}
	}
	fmt.Println(A - B)
}
