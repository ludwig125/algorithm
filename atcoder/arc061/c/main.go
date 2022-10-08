package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sum int

func main() {
	var s string
	fmt.Scan(&s)
	a := toSlice(s)
	p := make([]int, len(a)-1)

	if len(a) == 1 {
		fmt.Println(a[0])
		return
	}
	solve(a, 0, 1, &p)
	solve(a, 0, -1, &p)
	fmt.Println(sum)
}

func toSlice(s string) []int {
	var a []int
	for i := 0; i < len(s); i++ {
		a = append(a, toInt(string(s[i])))
	}
	return a
}

func solve(a []int, i int, b int, p *[]int) {
	if i == len(a)-1 {
		//fmt.Println(a, i, *p)
		return
	}
	(*p)[i] = b
	// fmt.Println(a, i, *p)
	if i == (len(a) - 2) {
		//fmt.Println(genNum(a, p), getSum(genNum(a, p)))
		sum += getSum(genNum(a, p))
	}
	solve(a, i+1, 1, p)
	solve(a, i+1, -1, p)
}

func genNum(a []int, p *[]int) string {
	var s string
	for j := 0; j < len(a)-1; j++ {
		s += fmt.Sprintf("%d", a[j])
		if (*p)[j] == 1 {
			s += "+"
		}
	}
	s += fmt.Sprintf("%d", a[len(a)-1])
	return s
}

func getSum(s string) int {
	sum := 0
	for _, v := range strings.Split(s, "+") {
		sum += toInt(v)
	}
	return sum
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
