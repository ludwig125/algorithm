package main

import (
	"fmt"
	"sort"
)

func main(){
	a := make([]int, 6)
	for i:=0; i<6; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort( sort.Reverse( sort.IntSlice(a)))
	fmt.Println(a[2])
}