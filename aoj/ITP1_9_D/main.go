package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s str
	fmt.Scan(&s)

	var q int
	fmt.Scan(&q)

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		sc.Scan()
		cmd := strings.Split(sc.Text(), " ")

		switch cmd[0] {
		case "replace":
			s.replace(toInt(cmd[1]), toInt(cmd[2]), cmd[3])
		case "reverse":
			s.reverse(toInt(cmd[1]), toInt(cmd[2]))
		case "print":
			s.print(toInt(cmd[1]), toInt(cmd[2]))
		}
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type str string

func (s *str) replace(a, b int, p string) {
	*s = str(string((*s)[:a]) + p + string((*s)[b+1:]))
}
func (s *str) reverse(a, b int) {
	var reversed string
	for i := b; i >= a; i-- {
		reversed += string((*s)[i])
	}
	*s = str(string((*s)[:a]) + reversed + string((*s)[b+1:]))
}
func (s str) print(a, b int) {
	fmt.Println(s[a : b+1])
}
