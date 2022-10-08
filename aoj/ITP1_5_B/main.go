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
			if h == 0 || h == H-1 {
				for w := 0; w < W; w++ {
					fmt.Print("#")
				}
				fmt.Println()
				continue
			}
			fmt.Print("#")
			for w := 1; w < W-1; w++ {
				fmt.Print(".")
			}
			fmt.Println("#")
		}
		fmt.Println()
	}

}
