package main

import (
	"fmt"
)

func main() {

	for {
		var s string
		fmt.Scan(&s)
		if s == "-" {
			break
		}
		var m int
		fmt.Scan(&m)

		var h int
		for i := 0; i < m; i++ {
			fmt.Scan(&h)
			s = shuffle(s, h)
		}
		fmt.Println(s)
	}
}

func shuffle(s string, m int) string {
	return s[m:] + s[:m]
}
