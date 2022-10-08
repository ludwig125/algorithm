package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// 奇数が1以外の奇数個あると、最終的に奇数と偶数が残るのでだめ
	// 奇数が偶数個あるかどうかわかればいい
	var odds []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 1 {
			odds = append(odds, num)
		}
	}
	if len(odds)%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
