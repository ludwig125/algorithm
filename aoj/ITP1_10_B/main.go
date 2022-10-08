package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	A := float64(a)
	B := float64(b)
	C := float64(c)

	S := A * math.Sin(C*math.Pi/180) * B / 2
	L := math.Sqrt(A*A+B*B-2*A*B*math.Cos(C*math.Pi/180)) + A + B
	fmt.Printf("%f\n", S)
	fmt.Printf("%f\n", L)
	fmt.Printf("%f\n", B*math.Sin(C*math.Pi/180))
}
