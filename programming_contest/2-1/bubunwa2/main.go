package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var k int
	fmt.Scan(&k)

	if solve(n, a, k, 0, 0) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func solve(n int, a []int, k int, i int, sum int) bool {
	if i == n {
		return sum == k
	}
	if solve(n, a, k, i+1, sum) {
		return true
	}
	if solve(n, a, k, i+1, sum+a[i]) {
		return true
	}
	return false
}
