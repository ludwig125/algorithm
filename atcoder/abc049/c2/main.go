package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for {
		if s == "" {
			fmt.Println("YES")
			break
		}
		if !strings.HasSuffix(s, "dream") && !strings.HasSuffix(s, "dreamer") && !strings.HasSuffix(s, "erase") && !strings.HasSuffix(s, "eraser") {
			fmt.Println("NO")
			break
		}

		for _, v := range []string{"dreamer", "dream", "eraser", "erase"} {
			s = cutSuffix(s, v)
		}
	}
}

func cutSuffix(s, target string) string {
	if strings.HasSuffix(s, target) {
		s = strings.TrimSuffix(s, target)
	}
	return s
}
