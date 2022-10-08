package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n-1; i++ {
		if a[i+1] == a[i] {
			fmt.Println("stay")
		} else if diff := a[i+1] - a[i]; diff > 0 {
			fmt.Println("up", diff)
		} else if diff := a[i+1] - a[i]; diff < 0 {
			fmt.Println("down", -1*diff)
		}
	}
}
