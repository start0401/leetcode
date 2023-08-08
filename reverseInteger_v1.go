package main

import (
	"log"
	"math"
	"strconv"
)

func main() {
	log.Printf("reverse %v ", reverse_v1(1534236469))
}

func reverse_v1(x int) int {
	z := x
	if x < 0 {
		z = -z
	}
	s := strconv.Itoa(z)
	y := ""
	for i := len(s) - 1; i >= 0; i-- {
		y += string(s[i])
	}
	num, _ := strconv.Atoi(y)
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
