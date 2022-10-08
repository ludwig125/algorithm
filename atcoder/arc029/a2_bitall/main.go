package main

import (
	"fmt"
	"sort"
)

func main() {
	var n uint
	fmt.Scan(&n)
	t := make([]int, n)
	for i := 0; uint(i) < n; i++ {
		fmt.Scan(&t[i])
	}
	sumList := check(n, t)
	sort.Ints(sumList)

	sumT := 0
	for _, v := range t {
		sumT += v
	}
	halfT := float64(sumT) / 2

	for _, v := range sumList {
		if float64(v) >= halfT {
			fmt.Println(v)
			return
		}
	}
}

func check(n uint, t []int) []int {
	var sumList []int
	for bit := 0; bit < 1<<n; bit++ {
		sum := 0
		for i := 0; uint(i) < n; i++ {
			if (bit & (1 << uint(i))) > 0 { // bitのi番目に1が立っているか？
				sum += t[i]
			}
		}
		sumList = append(sumList, sum)
	}
	return sumList
}
