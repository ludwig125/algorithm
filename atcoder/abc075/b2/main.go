package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := makeS(h, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(check(s, i, j, h, w))
		}
		fmt.Println()
	}
}

func makeS(h, w int) [][]string {
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
		var in string
		fmt.Scan(&in)
		for j := 0; j < w; j++ {
			s[i][j] = string(in[j])
		}
	}
	return s
}

func check(s [][]string, i, j, h, w int) string {
	if s[i][j] == "#" {
		return "#"
	}

	cnt := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if i+y == h || i+y == -1 || j+x == w || j+x == -1 || (x == 0 && y == 0) {
				continue
			}
			if s[i+y][j+x] == "#" {
				cnt++
			}
		}
	}
	return fmt.Sprintf("%d", cnt)
}
