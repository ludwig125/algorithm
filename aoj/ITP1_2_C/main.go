package main

import "fmt"

func main() {
	l := make([]int, 3)
	for i := range l {
		fmt.Scan(&l[i])
	}

	// if l[0] > l[1] {
	// 	l[0], l[1] = l[1], l[0]
	// }

	// if l[0] > l[2] {
	// 	l[0], l[2] = l[2], l[0]
	// }

	// if l[1] > l[2] {
	// 	l[1], l[2] = l[2], l[1]
	// }

	// fmt.Println(l[0], l[1], l[2])

	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] > l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
	fmt.Println(l[0], l[1], l[2])

}
