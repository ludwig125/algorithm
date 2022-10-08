package main

import (
	"fmt"
)

func main() {
	var m, f, r int
	for {
		fmt.Scan(&m)
		fmt.Scan(&f)
		fmt.Scan(&r)
		if m == -1 && f == -1 && r == -1 {
			break
		}
		judge(m, f, r)
	}
}

func judge(m, f, r int) {
	if m == -1 || f == -1 {
		fmt.Println("F")
		return
	}
	total := fix(m) + fix(f)
	if total >= 80 {
		fmt.Println("A")
	} else if total >= 65 && total < 80 {
		fmt.Println("B")
	} else if total >= 50 && total < 65 {
		fmt.Println("C")
	} else if total >= 30 && total < 50 {
		if r >= 50 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
		}
	} else if total < 30 {
		fmt.Println("F")
	}
}

func fix(n int) int {
	if n == -1 {
		return 0
	}
	return n
}
