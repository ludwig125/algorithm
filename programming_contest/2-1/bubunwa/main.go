package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)

	sort.Ints(a)
	//fmt.Println(n, k, a)
	if check(a, k, 0, 0) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func check(a []int, k int, i int, sum int) bool {
	if i == len(a) {
		fmt.Println("goal:", i, sum, k, sum == k)
		return sum == k
	}

	//sum += a[i]
	fmt.Println(i, a[i], sum)
	//check(a, k, i+1, sum)

	if check(a, k, i+1, sum) {
		fmt.Println("2:", i+1, sum)
		return true
	}
	// 今までのsumにa[i]を加えてｋに一致する場合はtrueが返る
	if check(a, k, i+1, sum+a[i]) {
		fmt.Println("1:", i+1, sum+a[i])
		return true
	}
	return false
}
