package main

import "fmt"

func main() {
	var x, y int
	for {
		fmt.Scan(&x)
		fmt.Scan(&y)

		if x == 0 && y == 0 {
			break
		}

		if x < y {
			fmt.Println(x, y)
		} else {
			fmt.Println(y, x)
		}
	}
}
