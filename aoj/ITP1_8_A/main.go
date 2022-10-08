package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var l string
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	l = s.Text()

	for i := 0; i < len(l); i++ {
		fmt.Print(string(conv(l[i])))
	}
	fmt.Println()

	// 面白い回答
	// for _, c := range(scanner.Text()) {
	// 	if c > 0x40 && (c-0x41)%0x20 < 0x1a {
	// 		c ^= 0x20
	// 	}
	// 	s += string(c)
	// }
	// fmt.Println(s)
}

func conv(n byte) byte {
	if n >= 65 && n <= 90 { // A to Z
		return n + 32 // A(65) から Z(90) の場合は a(97) から z(122) に変換
	}
	if n >= 97 && n <= 122 { // a to z
		return n - 32
	}
	return n
}
