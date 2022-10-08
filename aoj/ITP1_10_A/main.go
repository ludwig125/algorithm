package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	xlen := math.Abs(x2 - x1)
	ylen := math.Abs(y2 - y1)
	fmt.Printf("%f\n", math.Sqrt(xlen*xlen+ylen*ylen))
}
