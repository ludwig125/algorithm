package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	allMn := genMarkNum()

	var N int
	fmt.Scan(&N)

	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < N; i++ {
		s.Scan()
		allMn = filter(allMn, s.Text())
	}
	for _, v := range allMn {
		fmt.Println(v)
	}
}

func genMarkNum() []string {
	var s []string
	for _, m := range []string{"S", "H", "C", "D"} {
		for i := 1; i <= 13; i++ {
			s = append(s, fmt.Sprintf("%s %d", m, i))
		}
	}
	return s
}

func filter(mn []string, target string) []string {
	var filteredMn []string
	for _, v := range mn {
		if v == target {
			continue
		}
		filteredMn = append(filteredMn, v)
	}
	return filteredMn
}
