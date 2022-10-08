package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(i * 2)
}
