package main

import (
	"fmt"
)

func main() {
	var st string
	fmt.Scan(&st)

	sum := 0
	for _, i := range st {
		if i == '1' {
			sum += 1
		}
	}
	fmt.Println(sum)

}
