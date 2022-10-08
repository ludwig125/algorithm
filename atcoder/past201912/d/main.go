package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(run(a))
}

func run(a []int) string {
	sort.Sort(sort.IntSlice(a))
	if double, non := doubleAndNon(a); double > 0 || non > 0 {
		return fmt.Sprintf("%d %d", double, non)
	}
	return "Correct"
}

func doubleAndNon(a []int) (int, int) {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}

	var double, non int
	for i := 1; i <= len(a); i++ {
		if m[i] == 2 {
			double = i
		}
		if m[i] == 0 {
			non = i
		}
	}

	return double, non
}
