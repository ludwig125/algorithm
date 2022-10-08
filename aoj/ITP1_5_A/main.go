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
		for h := 0; h < H; h++ {
			for w := 0; w < W; w++ {
				fmt.Print("#")
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
