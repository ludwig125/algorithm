package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	sc.Scan()
	p := sc.Text()

	check(s, p)

}
func check(s, p string) {
	ret := "No"
	for i := 0; i < len(s); i++ {
		newStr := part(s, i, len(s)) + part(s, 0, i)
		if sContainsP(newStr, p) {
			ret = "Yes"
			break
		}
	}
	fmt.Println(ret)
}

func part(s string, b, e int) string {
	var p string
	for i := b; i < e; i++ {
		p += string(s[i])
	}
	return p
}

func sContainsP(s, p string) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != s[i] {
			return false
		}
	}
	return true
}
