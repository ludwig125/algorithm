package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)

	fmt.Printf("%.6f %.6f\n", r*r*math.Pi, r*2*math.Pi)
}
