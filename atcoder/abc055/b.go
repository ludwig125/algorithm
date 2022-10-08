package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var power int64
	power = 1
	for i := 1; i <= N; i++ {
		power *= int64(i)
		power %= 1000000007
	}
	fmt.Println(power)

}
