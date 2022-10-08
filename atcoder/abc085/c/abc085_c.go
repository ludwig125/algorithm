// https://qiita.com/drken/items/fd4e5e3630d0f5859067
package main

import (
	"fmt"
	//"sort"
)

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)

	//	check := func() {
	//		for a := 0; a <= N; a++ {
	//			for b := 0; b <= N; b++ {
	//				for c := 0; c <= N; c++ {
	//					if (a*10000+b*5000+c*1000) == Y && (a+b+c) == N {
	//						fmt.Printf("%d %d %d %d\n", a, b, c, a*10000+b*5000+c*1000)
	//						return
	//					}
	//				}
	//			}
	//		}
	//		fmt.Println("-1 -1 -1")
	//	}
	check := func() {
		for a := 0; a <= N; a++ {
			for b := 0; b <= (N - a); b++ {
				c := N - a - b
				if (a*10000 + b*5000 + c*1000) == Y {
					//fmt.Printf("%d %d %d %d\n", a, b, c, a*10000+b*5000+c*1000)
					fmt.Println(a, b, c)
					return
				}
			}
		}
		fmt.Println("-1 -1 -1")
	}
	check()

}
