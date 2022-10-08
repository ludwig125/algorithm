package main

import "fmt"

func main() {
	var W, H, x, y, r int
	fmt.Scan(&W)
	fmt.Scan(&H)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&r)

	if 0 <= (x-r) && (x+r) <= W && 0 <= (y-r) && (y+r) <= H {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
