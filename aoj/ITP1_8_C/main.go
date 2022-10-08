package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	m := newAlpha()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()
		if l == "" {
			break
		}
		countAlpha(&m, l)
	}
	printAlpha(m)

}

func newAlpha() map[string]int {
	m := make(map[string]int)
	for i := 97; i <= 122; i++ {
		m[string(i)] = 0
	}
	return m
}

func countAlpha(m *map[string]int, l string) {
	for i := range l {
		a := strings.ToLower(string(l[i]))
		if a >= "a" && a <= "z" {
			(*m)[a]++
		}
	}
}

func printAlpha(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s : %d\n", k, m[k])
	}
}
