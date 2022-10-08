package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	min := 1000000
	max := -1000000
	sum := 0
	for i := range a {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
		sum += a[i]
	}
	fmt.Println(min, max, sum)
}
