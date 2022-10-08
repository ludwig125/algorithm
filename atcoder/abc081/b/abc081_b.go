package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	Al := make([]int, N)
	for i := range Al {
		fmt.Scan(&Al[i])
	}

	isEven := func(x int) int {
		if x%2 == 0 {
			return 0
		} else {
			return 1
		}
	}

	areAllEven := func(l []int) bool {
		evenCount := 0
		for _, v := range l {
			evenCount += isEven(v)
		}
		if evenCount == 0 {
			return true
		} else {
			return false
		}
	}

	divide2 := func(l []int) []int {
		var nl []int
		for _, v := range l {
			nl = append(nl, v/2)
		}
		return nl
	}

	divedeNum := 0
	l := Al
	for {
		if areAllEven(l) == true {
			l = divide2(l)
			divedeNum += 1
		} else {
			break
		}
	}

	fmt.Println(divedeNum)
}
