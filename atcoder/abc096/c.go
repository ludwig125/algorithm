package main

import (
	"fmt"
	"strings"
)

func isNeighborHash(s [][]string, i int, j int, h int, w int) bool {
	if i-1 > 0 {
		if s[i-1][j] == "#" {
			return true
		}
	}
	if i+1 < h {
		if s[i+1][j] == "#" {
			return true
		}
	}
	if j-1 > 0 {
		if s[i][j-1] == "#" {
			return true
		}
	}
	if j+1 < w {
		if s[i][j+1] == "#" {
			return true
		}
	}
	return false
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var s [][]string
	for i := 0; i < h; i++ {
		var l string
		fmt.Scan(&l)
		s = append(s, strings.Split(l, ""))
	}
	//	for i := 0; i < h; i++ {
	//		for j := 0; j < w; j++ {
	//			fmt.Print(s[i][j])
	//		}
	//		fmt.Println()
	//	}

	check := func() string {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				//fmt.Println("--------")
				//fmt.Print(s[i][j])

				if s[i][j] == "#" {
					//fmt.Println(i, j, s[i][j], isNeighborHash(s, i, j, h, w))
					//fmt.Println(i, j, s[i][j], s[i-1][j], s[i+1][j], s[i][j-1], s[i][j+1])
					if !isNeighborHash(s, i, j, h, w) {
						return "No"
					}
				}
			}
		}
		return "Yes"
	}
	fmt.Println(check())
}
