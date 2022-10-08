package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	diffAc := (a[2] - a[0])
	diffBc := (a[2] - a[1])
	count := diffAc/2 + diffBc/2

	// それぞれの差分を２で割ったときのあまりを見る
	// <diffAc % 2, diffBc % 2> == <0, 0> => これ以上加算不要
	// <diffAc % 2, diffBc % 2> == <0, 1> => ２つに１足して１つに２足す処理が必要(2step必要)
	// <diffAc % 2, diffBc % 2> == <1, 0> => 同上
	// <diffAc % 2, diffBc % 2> == <1, 1> => ２つに１足す処理が必要(1step必要)

	fmt.Printf("a->c %d %d, b->c %d %d\n", diffAc/2, diffAc%2, diffBc/2, diffBc%2)

	if (diffAc%2 + diffBc%2) == 1 {
		count += 2
	} else if (diffAc%2 + diffBc%2) == 2 {
		count += 1
	}
	fmt.Println(count)
}
