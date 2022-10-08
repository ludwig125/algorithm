package main

import (
	"fmt"
)

func main() {
	var n, x int
	for {
		fmt.Scan(&n)
		fmt.Scan(&x)
		if n == 0 && x == 0 {
			break
		}
		calc(n, x)
	}
}

func calc(n, x int) {
	l := makeNList(n)
	fmt.Println(check(l, n, x))

}

func makeNList(n int) []int {
	var l []int
	for i := 1; i <= n; i++ {
		l = append(l, i)
	}
	return l
}

func check(l []int, n, x int) int {
	cnt := 0
	for k := 2; k < n; k++ {
		for j := 1; j < k; j++ {
			for i := 0; i < j; i++ {
				if l[i]+l[j]+l[k] == x {
					cnt++
				}
			}
		}
	}
	return cnt
}
