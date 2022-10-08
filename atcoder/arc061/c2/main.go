package main

import (
	"fmt"
	"strconv"
)

// 参考 https://atcoder.jp/contests/arc061/submissions/10194970
// 途中までしかやっていない
//var num [4][4]int

func main() {
	//s := "1256"
	s := "125"
	sum := dfs(s, 0, 0, 0)
	fmt.Println(sum)
}

func dfs(s string, i, j int, sum int) int {
	fmt.Println("--")
	if i == len(s) {
		if i == j {
			return sum
		}
		return 0
	}
	n, _ := strconv.Atoi(s[j : i+1])
	fmt.Printf("i: %d, j: %d, s[%d:%d]=%d\n", i, j, j, i+1, n)
	sum1 := dfs(s, i+1, j, sum)
	sum2 := dfs(s, i+1, i+1, sum+n)
	//fmt.Println("sum1", sum1, "sum2", sum2)
	return sum1 + sum2
}

// var sum int

// func main() {
// 	//s := "1256"
// 	s := "125"
// 	dfs(s, 0, 0)

// 	fmt.Println(sum)
// }

// func dfs(s string, i, j int) {
// 	fmt.Println("--")
// 	if i == len(s) {
// 		return
// 	}
// 	n, _ := strconv.Atoi(s[j : i+1])
// 	fmt.Printf("i: %d, j: %d, s[%d:%d]=%d\n", i, j, j, i+1, n)
// 	fmt.Println(sum, n)
// 	dfs(s, i+1, j)
// 	dfs(s, i+1, i+1)
// 	return
// }

func fn(s string, i, j int) int {
	n, _ := strconv.Atoi(s[j:i])
	fmt.Printf("i: %d, j: %d, s[%d:%d]=%d\n", i, j, j, i, n)
	return n
}

// func dfs(s string, i, j int) int {
// 	if i == len(s) {
// 		return 0
// 	}
// 	//fmt.Printf("s[%d:%d]=%s\n", j, i+1, s[j:i+1])
// 	n, _ := strconv.Atoi(s[j : i+1])
// 	// fmt.Println(j, i+1, "n:", n)

// 	// n1 := dfs(s, i+1, j)
// 	// n2 := dfs(s, i+1, j+1)

// 	_ = dfs(s, i+1, j)
// 	//_ = dfs(s, i+1, i+1)
// 	//sum = n1 + n2
// 	//fmt.Println("n1:", n1, "n2", n2, n)
// 	fmt.Println(j, i, "n:", n)
// 	return n
// }

// func dfs(s string, i int, lasti int, sum int) int {
// 	if i == len(s) {
// 		if lasti == i {
// 			return sum
// 		}
// 		return 0
// 	}
// 	fmt.Println(i, lasti, s[lasti:i+1])
// 	n, _ := strconv.Atoi(s[lasti : i+1])
// 	sum1 := dfs(s, i+1, lasti, sum)
// 	sum2 := dfs(s, i+1, i+1, sum+n)

// 	return sum1 + sum2
// }

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	sum := dfs(s, 0, 0, 0)
// 	fmt.Println(sum)
// }

// package main

// import "fmt"

// func main() {
// 	num := 5
// 	p := make([]int, num)
// 	dfs(num, -1, 0, &p)
// }

// func dfs(num, i, flg int, p *[]int) {
// 	if i == num {
// 		return
// 	}
// 	if i == (num - 1) {
// 		fmt.Printf("i: %d, p: %v\n", i, *p)

// 	}

// 	if i >= 0 {
// 		(*p)[i] = flg
// 	}
// 	dfs(num, i+1, 1, p)
// 	dfs(num, i+1, -1, p)
// }
