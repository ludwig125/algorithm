package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)

	s := bufio.NewScanner(os.Stdin)
	cnt := 0
	for s.Scan() {
		l := s.Text()
		if l == "END_OF_TEXT" {
			break
		}
		cnt += find(l, w)
	}
	fmt.Println(cnt)
}

func find(l, w string) int {
	cnt := 0
	for _, each := range strings.Split(l, " ") {
		if strings.ToLower(each) == w {
			cnt++
		}
	}
	return cnt
}
