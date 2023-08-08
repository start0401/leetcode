package main

import (
	"log"
)

func main() {
	log.Printf("myPow %v", myPow(2, 10))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	y := myPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}
