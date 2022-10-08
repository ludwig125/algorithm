package main

import "fmt"

func sumEachOrder(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
		//	if n <= 0 {
		//		break
		//	}
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
