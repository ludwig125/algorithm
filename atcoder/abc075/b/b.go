package main

import (
	"fmt"
	"strings"
)

func isHash(s string) int {
	if s == "#" {
		return 1
	}
	return 0
}

// 前後左右斜めが＃かどうか判定する関数
func checkAround(p [][]string, i int, j int, h int, w int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			//	fmt.Printf("i %d j %d k %d l %d \n", i, j, k, l)
			if k >= 0 && k < h && l >= 0 && l < w {
				if !(k == i && l == j) {
					//			fmt.Printf("i %d j %d k %d l %d p[k][l] %v\n", i, j, k, l, p[k][l])
					//fmt.Printf("i %d j %d k %d l %d\n", i, j, k, l)
					count += isHash(p[k][l])
				}
			}
			//fmt.Printf("i %d j %d k %d l %d isHash %v\n", i, j, k, l, isHash(p[k][l]))
		}
	}
	return count
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var p [][]string
	for i := 0; i < h; i++ {
		var l string
		fmt.Scan(&l)
		p = append(p, strings.Split(l, ""))
	}
	//	for i := 0; i < h; i++ {
	//		for j := 0; j < w; j++ {
	//			fmt.Print(p[i][j])
	//		}
	//		fmt.Println()
	//	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if isHash(p[i][j]) == 1 {
				fmt.Print("#")
			} else {
				//fmt.Printf("i %d j %d\n", i, j)
				fmt.Print(checkAround(p, i, j, h, w))
				//checkAround(p, i, j, h, w)
				//fmt.Print(p[i][j])
			}
		}
		fmt.Println()
	}
	//fmt.Println(checkAround(p, 2, 3, h, w))
}
