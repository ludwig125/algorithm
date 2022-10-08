package main

import "fmt"

func main() {
	var a, b int
	var op string
L:
	for {
		fmt.Scan(&a)
		fmt.Scan(&op)
		fmt.Scan(&b)

		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "?":
			break L
		}
	}
}
