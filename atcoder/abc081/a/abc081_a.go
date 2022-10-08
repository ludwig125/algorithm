package main

import (
	"fmt"
	"strconv"
)

func main() {
	var st string
	fmt.Scan(&st)

	sum := 0
	for _, i := range st {
		// string 型に一旦する方法
		j := fmt.Sprintf("%c", i)
		k, _ := strconv.Atoi(j)
		sum += k
	}
	fmt.Println(sum)

}
