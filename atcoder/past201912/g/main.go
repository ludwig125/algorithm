package main

import (
	"fmt"
	"sort"
)

// 途中でやめた

func main() {
	var n int
	fmt.Scan(&n)

	a := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = make([]int, n-1-i)
		for j := 0; j < n-1-i; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	var score []int
	for i, v := range a {
		for j, v2 := range v {
			fmt.Println(i, j, v2)
			score = append(score, v2)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	fmt.Println(score)
}

// func check(i int, sum int){
// 	if i ==
// }
