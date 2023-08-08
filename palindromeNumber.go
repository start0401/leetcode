package main

import (
	"log"
)

func main() {
	log.Printf("reverse %v ", isPalindrome(0))
}

func isPalindrome(x int) bool {
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

	if x >= 0 && num == x {
		return true
	}
	return false
}
