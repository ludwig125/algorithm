package main

import (
	"fmt"
	"strconv"
)

func main() {
	//a := []int{1, 2, 2, 2}
	//a := []int{0, 2, 9, 0}
	var s string
	fmt.Scan(&s)
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		a[i] = n
	}
	//fmt.Println(a, s)
	fmt.Println(solve(a, 0, a[0], fmt.Sprintf("%d", a[0]), ""))
}

func solve(a []int, i, ans int, f, op string) string {
	if i == 4 {
		//fmt.Println(f, ans)
		if ans == 7 {
			return fmt.Sprintf("%s=%d", f, ans)
		}
		return ""
	}

	if op == "+" {
		ans += a[i]
		f += fmt.Sprintf("+%d", a[i])

	} else if op == "-" {
		ans -= a[i]
		f += fmt.Sprintf("-%d", a[i])
	}
	if res := solve(a, i+1, ans, f, "+"); res != "" {
		return res
	}
	if res := solve(a, i+1, ans, f, "-"); res != "" {
		return res
	}
	return ""
}

// func solve(a []int, i, ans int, f, op string) {
// 	if i == 4 {
// 		if ans==7{
// 			fmt.Println(f,ans)
// 			return f+"="+ans
// 		}
// 		return ""
// 	}

// 	if op == "+" {
// 		ans += a[i]
// 		f += fmt.Sprintf("+%d", a[i])

// 	} else {
// 		ans -= a[i]
// 		f += fmt.Sprintf("-%d", a[i])
// 	}
// 	solve(a, i+1, ans, f, "+")
// 	solve(a, i+1, ans, f, "-")
// 	return
// }
