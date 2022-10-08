package main

import (
	"fmt"
)

func main() {
	var H, W int
	for {
		fmt.Scan(&H)
		fmt.Scan(&W)

		if H == 0 && W == 0 {
			break
		}

		m := mark{flg: 1}
		for h := 0; h < H; h++ {
			if h%2 == 0 {
				m.flg = 1
			} else {
				m.flg = -1
			}
			for w := 0; w < W; w++ {
				m.draw()
			}
			fmt.Println()
		}
		fmt.Println()

		// 良い解法
		// for i := 0; i < H; i++ {
		// 	for j := 0; j < W; j++ {
		// 		if (i+j)%2 == 0 {
		// 			fmt.Print("#")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println("")
		// }
	}
}

type mark struct {
	flg int
}

func (m *mark) draw() {
	if m.flg == 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	m.flg *= -1
}
