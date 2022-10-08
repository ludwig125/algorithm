package main

import "fmt"

func sumEachOrder(n int) int {
	sum := 0
	for f := n; f > 0; f /= 10 {
		sum += f % 10
	}
	return sum
}

func main() {
	var N, A, B int
	fmt.Scan(&N)
	fmt.Scan(&A)
	fmt.Scan(&B)

	ans := 0
	for n := 0; n <= N; n++ {
		s := sumEachOrder(n)
		if s >= A && s <= B {
			ans += n
		}
	}
	fmt.Println(ans)
}
