package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var startX, startY, goalX, goalY int
	var line string
	c := make([][]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&line)
		c[i] = make([]string, w)
		for j := 0; j < w; j++ {
			c[i][j] = string(line[j])
			if c[i][j] == "s" {
				startX = j
				startY = i
			}
			if c[i][j] == "g" {
				goalX = j
				goalY = i
			}
		}
	}

	reached := make([][]int, h)
	for i := 0; i < h; i++ {
		reached[i] = make([]int, w)
	}

	search(h, w, c, &reached, startX, startY)

	// fmt.Println(h, w)
	// fmt.Println(startX, startY)
	// fmt.Println(goalX, goalY)
	// for _, v := range c {
	// 	fmt.Println(v)
	// }
	// for _, v := range reached {
	// 	fmt.Println(v)
	// }
	if reached[goalY][goalX] == 1 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func search(h, w int, c [][]string, r *[][]int, x, y int) {
	if x < 0 || y < 0 || x >= w || y >= h || c[y][x] == "#" {
		return
	}

	if (*r)[y][x] == 1 { // 以前訪れたことがあるなら戻る
		return
	}
	(*r)[y][x] = 1

	search(h, w, c, r, x+1, y)
	search(h, w, c, r, x-1, y)
	search(h, w, c, r, x, y+1)
	search(h, w, c, r, x, y-1)
}
