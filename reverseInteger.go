package main

import (
	"log"
	"math"
)

func main() {
	log.Printf("reverse %v ", reverse(120))
}

func reverse(x int) int {
	z := x
	if x < 0 {
		z = -z
	}
	num := 0
	for z > 0 {
		remainder := z % 10
		num *= 10
		num += remainder
		z /= 10
	}

	if x < 0 {
		num = -num
	}

	if num > math.MaxInt32-1 {
		return 0
	}

	if num < math.MinInt32 {
		return 0
	}

	return num
}
