package main

import (
	"fmt"
)

func main() {
	var r, c int
	fmt.Scan(&r)
	fmt.Scan(&c)

	s := makeEmptySheet(r, c)
	s = calcSheet(s, r, c)
	printSheet(s)
}

func makeEmptySheet(r, c int) [][]int {
	s := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		s[i] = make([]int, c+1)
	}
	return s
}

func calcSheet(s [][]int, r, c int) [][]int {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			var num int
			fmt.Scan(&num)
			s[i][j] = num
			s[i][c] += num
			s[r][j] += num
		}
	}

	for j := 0; j < c; j++ {
		s[r][c] += s[r][j]
	}
	return s
}

func printSheet(s [][]int) {
	for i := range s {
		for j := range s[i] {
			fmt.Printf("%d ", s[i][j])
		}
		fmt.Println()
	}
}
