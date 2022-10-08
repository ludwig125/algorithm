package main

import (
	"fmt"
	"sort"
)

func main() {
	//var a, b, c, d int
	//fmt.Scan(&a, &b, &c, &d)

	l := make([]int, 4)
	for i := range l {
		fmt.Scan(&l[i])
	}
	a := l[0]
	b := l[1]
	c := l[2]
	d := l[3]
	sort.Ints(l)

	if l[1] >= a && l[1] >= c && l[2] <= b && l[2] <= d {
		fmt.Println(l[2] - l[1])
	} else {
		fmt.Println(0)
	}
}
