package main

import (
	"fmt"
)

var i = 1

func main() {
	var n int
	fmt.Scan(&n)

	call(n)
}

func call(n int) {
	for {
		x := i
		if checkNum(x) {
			include3(x)
		}
		if endCheckNum(n) {
			break
		}
	}
	fmt.Println()

}

func checkNum(x int) bool {
	//fmt.Println("checkNum", i)
	if x%3 == 0 {
		fmt.Printf(" %d", i)
		return false
	}
	return true
}

func include3(x int) {
	//time.Sleep(1 * time.Second)
	//fmt.Println("include3", i)
	if x%10 == 3 {
		fmt.Printf(" %d", i)
		return
	}
	if x/10 != 0 {
		include3(x / 10)
	}
}

func endCheckNum(n int) bool {
	//fmt.Println("endCheckNum", i)
	i++
	if i <= n {
		return false
	}
	return true
}
