package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		s := make([]int, n)
		for i := range s {
			fmt.Scan(&s[i])
		}
		fmt.Printf("%f\n", standardDevidation(s))
	}
}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func average(s []int) float64 {
	return float64(sum(s)) / float64(len(s))
}

func standardDevidation(s []int) float64 {
	sd := 0.0
	for _, v := range s {
		sd += (float64(v) - average(s)) * (float64(v) - average(s))
	}
	return math.Sqrt(sd / float64(len(s)))
}
