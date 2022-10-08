package main

import (
	"fmt"
)

func hasStr(S string, str string) bool {
	return len(str) <= len(S) && S[len(S)-len(str):] == str
}

func main() {
	var S string
	fmt.Scan(&S)

	list := []string{"dream", "dreamer", "erase", "eraser"}
	for {
		tmp := S
		for _, v := range list {
			if hasStr(S, v) {
				S = S[:len(S)-len(v)]
			}
		}
		if tmp == S {
			break
		}
	}
	if len(S) == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
