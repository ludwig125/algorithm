package main

import "fmt"

func count(A int, B int, C int, X int) int {
	count := 0
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				if (a*500 + b*100 + c*50) == X {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	var A, B, C, X int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&X)
	fmt.Println(count(A, B, C, X))
}
