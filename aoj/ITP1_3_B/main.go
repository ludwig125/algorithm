package main

import "fmt"

func main() {
	var x int
	i := 0
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		i++
		fmt.Printf("Case %d: %d\n", i, x)
	}
}
