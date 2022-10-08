package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	for {
		fmt.Scan(&n)
		if len(n) == 1 && toInt(n) == 0 {
			break
		}
		fmt.Println(sum(n))
	}
}

func sum(n string) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += toInt(string(n[i]))
	}
	return sum
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
