package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	k := make([]int, n)
	for i := range k {
		fmt.Scan(&k[i])
	}

	fmt.Println(check(n, m, k))
}

func check(n, m int, k []int) string {
	for i := 0; i < n; i++ {
		for i2 := 0; i2 < n; i2++ {
			for i3 := 0; i3 < n; i3++ {
				for i4 := 0; i4 < n; i4++ {
					if k[i]+k[i2]+k[i3]+k[i4] == m {
						fmt.Println("Yes", k[i], k[i2], k[i3], k[i4])
						return "Yes"
					}
				}
			}
		}
	}
	return "No"
}
