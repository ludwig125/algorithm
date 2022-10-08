package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	var chosenTList []int
	m := dfs(n, t, 0, 0, chosenTList)

	sumT := 0
	for _, v := range t {
		sumT += v
	}
	halfT := float64(sumT) / 2

	for _, t := range sortT(m) {
		if float64(t) >= halfT {
			fmt.Println(t)
			return
		}
	}
}

func dfs(n int, t []int, i, sum int, chosen []int) map[int][]int {
	if i == n {
		return map[int][]int{sum: chosen}
	}
	m1 := dfs(n, t, i+1, sum, chosen) // tiを選ぶ
	chosen = append(chosen, t[i])
	m2 := dfs(n, t, i+1, sum+t[i], chosen) // tiを選ばない
	return merge(m1, m2)
}
func merge(m1, m2 map[int][]int) map[int][]int {
	res := map[int][]int{}

	for k, v := range m1 {
		res[k] = v
	}
	for k, v := range m2 {
		res[k] = v
	}
	return res
}

func sortT(m map[int][]int) []int {
	var a []int
	for k := range m {
		a = append(a, k)
	}
	sort.Ints(a)
	return a
}
