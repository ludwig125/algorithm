package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}

	y := make([]int, n)
	for i := range y {
		fmt.Scan(&y[i])
	}

	p1 := 0.0
	p2 := 0.0
	p3 := 0.0
	max := 0.0
	for i := 0; i < n; i++ {
		p1 += absDistance(x[i], y[i])
		p2 += absDistance(x[i], y[i]) * absDistance(x[i], y[i])
		p3 += absDistance(x[i], y[i]) * absDistance(x[i], y[i]) * absDistance(x[i], y[i])
		if max < absDistance(x[i], y[i]) {
			max = absDistance(x[i], y[i])
		}
	}
	fmt.Printf("%f\n", p1)
	fmt.Printf("%f\n", math.Sqrt(p2))
	fmt.Printf("%f\n", math.Cbrt(p3))
	fmt.Printf("%f\n", max)
}

func absDistance(x, y int) float64 {
	return math.Abs(float64(x) - float64(y))
}
