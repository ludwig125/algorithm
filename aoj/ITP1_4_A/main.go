package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("%d %d %.5f\n", int(a/b), a%b, float64(a)/float64(b))
}
